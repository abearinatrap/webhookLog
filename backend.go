package webhookLog

import (
	"fmt"
	"time"
)

func (l *DefaultLogger) runBack() {
	l.queue = make(chan string, l.queueSize)
	message := ""
	fmt.Printf("%s%s", logPrefix, "Backend logger running")
	waitTime := 1000
	for {
		//bounded execution time for log sending
		select {
		case newMessage := <-l.queue:
			message += string(newMessage)
			message, _ = l.sendMessage(message)
		case <-time.After(time.Duration(waitTime) * time.Millisecond):
			if len(message) > 0 {
				var statusCode int
				message, statusCode = l.sendMessage(message)
				if statusCode == 200 {
					waitTime = 60
				} else {
					waitTime += 20
				}
			} else {
				waitTime = 1000
			}
		}

	}
}
