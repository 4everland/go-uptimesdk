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
err := c.ReportWithOpt(message.Message{
        Value: 66,
        Values: []message.Quota{
            message.Quota{Name:"test1", Value: 64}
        },
    })
//with options
err := c.ReportWithOpt(
    message.Message{
        Value: 65,
    }, 
    "", 
    WithServerName("test-2"))
```