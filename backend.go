package webhookLog

import "fmt"

func (l *DefaultLogger) runBack() {
	l.queue = make(chan string, l.queueSize)
	message := ""
	fmt.Printf("Backend logger running")
	for {
		newMessage := <-l.queue
		message += string(newMessage)
		message, _ = l.sendMessage(message)
	}
}
