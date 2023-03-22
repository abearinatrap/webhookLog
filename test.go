package webhookLog

import (
	"fmt"
	"time"
)

var Log *DefaultLogger

func main() {
	// test logging
	Log = NewDefaultLogger("LOGGER_NAME", "WEBHOOK_KEY")
	start := time.Now()
	Log.Infof("Testing log")
	duration := time.Since(start)
	fmt.Printf("Time to send %v\n", duration)
}
