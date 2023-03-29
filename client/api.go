package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/4everland/go-uptimesdk/message"
	"io"
	"net/http"
	"time"
)

const (
	reportPath = "/uptime/v2/report"
)

type Client struct {
	opt    ReportOption
	client *http.Client
}

type ReportOption struct {
	Host          string
	Authorization string
	Header        message.Header
}

type Option func(*ReportOption)

func New(opts ...Option) *Client {
	opt := ReportOption{}
	for _, o := range opts {
		o(&opt)
	}

	return &Client{
		opt: opt,
		client: &http.Client{
			Timeout: time.Second * 3,
		},
	}
}

func (c *Client) ReportMany(m map[string]message.Message) error {
	return c.ReportWithExtra(m, "")
}

func (c *Client) Report(name string, m message.Message) error {
	return c.ReportMany(map[string]message.Message{name: m})
}

func (c *Client) ReportWithExtra(m map[string]message.Message, extra string) error {
	return c.ReportWithOpt(m, extra)
}

func (c *Client) ReportWithOpt(m map[string]message.Message, extra string, opts ...Option) error {
	opt := c.opt
	for _, o := range opts {
		o(&opt)
	}

	url := opt.Host + reportPath
	body := message.Package{
		ResourceMap: m,

		Header: opt.Header,
		Extra:  extra,
	}
	buf, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(buf))

	req.Header.Set("Authorization", opt.Authorization)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		msg, _ := io.ReadAll(resp.Body)
		errMsg := ""
		if msg != nil {
			errMsg = string(msg)
		}
		return fmt.Errorf("error reps code: %d, message: %s", resp.StatusCode, errMsg)
	}
	return nil
}
