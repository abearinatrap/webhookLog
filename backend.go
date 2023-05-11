package webhookLog

import (
	"fmt"
	"time"
)

func (l *DefaultLogger) runBack() {
	l.queue = make(chan string, l.queueSize)
	message := ""
	fmt.Printf("Backend logger running")
	for {
		//bounded execution time for log sending
		select {
		case newMessage := <-l.queue:
			message += string(newMessage)
			message, _ = l.sendMessage(message)
		case <-time.After(1 * time.Second):
			message, _ = l.sendMessage(message)
		}

	}
}
