# webhookLog
A simple package for using discord webhooks as a logging mechanism


### Usage
```
var Log *webhookLog.DefaultLogger
Log = webhookLog.NewDefaultLogger("LOGGER_NAME", "WEBHOOK_KEY")
Log.SetLevel(webhookLog.Debug)
Log.Info("Logging setup")
```

###### Install
`go get github.com/abearinatrap/webhookLog`