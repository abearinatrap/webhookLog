package webhookLog

import (
	"fmt"
	"time"
)

func (l *DefaultLogger) runBack() {
	l.queue = make(chan string, l.queueSize)
	message := ""
	fmt.Printf("Backend logger running")
	waitTime := 1000
	for {
		//bounded execution time for log sending
		select {
		case newMessage := <-l.queue:
			message += string(newMessage)
			message, _ = l.sendMessage(message)
		case <-time.After(time.Duration(waitTime) * time.Millisecond):
			if len(message) > 0 {
				message, _ = l.sendMessage(message)
				waitTime = 250
			} else {
				waitTime = 1000
			}
		}

	}
}
