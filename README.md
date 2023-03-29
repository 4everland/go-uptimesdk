# go-uptimesdk

Internal used client
### install
```
go get github.com/4everland/go-uptimesdk
```

### Usage

```go  
client := NewClient("http://127.0.0.1", "test", "test-server", "test-instance")
//simple report 
err := c.Report("test", message.Message{
		Value: 32,
	})
//with options
	err := c.ReportWithOpt(map[string]message.Message{
		"test": {
			Value: 32,
		},
	}, "", WithServerName("test-2"))
```