package gouptimesdk

import "github.com/4everland/go-uptimesdk/client"

func WithAuthorization(auth string) client.Option {
	return func(option *client.ReportOption) {
		option.Authorization = auth
	}
}

func WithHost(host string) client.Option {
	return func(option *client.ReportOption) {
		option.Host = host
	}
}

func WithServerName(serverName string) client.Option {
	return func(option *client.ReportOption) {
		option.Header.ServerName = serverName
	}
}

func WithInstanceName(instanceName string) client.Option {
	return func(option *client.ReportOption) {
		option.Header.InstanceName = instanceName
	}
}

func NewClient(host, auth, server, instance string) *client.Client {
	return client.New(WithHost(host), WithAuthorization(auth), WithServerName(server), WithInstanceName(instance))
}
