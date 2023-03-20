package gouptimesdk

import (
	"github.com/4everland/go-uptimesdk/message"
	"testing"
)

func TestReport(t *testing.T) {
	c := NewClient("http://127.0.0.1", "test", "test-server", "test-instance")
	err := c.ReportWithOpt(message.Message{
		Value: 32,
	}, "", WithServerName("test-2"))
	if err != nil {
		t.Fatalf("%s", err)
	}
}
