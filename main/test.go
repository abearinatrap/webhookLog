package main

import (
	"time"

	"github.com/abearinatrap/webhookLog"
)

var Log *webhookLog.DefaultLogger

func main() {
	// test logging
	Log = webhookLog.NewDefaultLogger("LOGGER_NAME", "WEBHOOK_KEY")
	Log.SetLevel(webhookLog.Debug)
	start := time.Now()
	Log.Infof("%v", start)
	Log.Infof("Testing")
}
