# webhookLog
A simple package for using discord webhooks as a logging mechanism

Rate limit is 16-17 logs a second, at 2000 characters max per webhook, this is 32000-34000 characters a second max you can log to discord

### Usage
```
var Log *webhookLog.DefaultLogger
Log = webhookLog.NewDefaultLogger("LOGGER_NAME", "WEBHOOK_KEY")
Log.SetLevel(webhookLog.Debug)
Log.Info("Logging setup")
```

###### Install
`go get github.com/abearinatrap/webhookLog@v0.1.5`