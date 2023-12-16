package webhookLog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const logPrefix = "WebhookLog: "

func log(s string) {
	fmt.Printf("%s%s", logPrefix, s)
}

func logf(s string, args ...interface{}) {
	message := fmt.Sprintf(s, args...)
	fmt.Printf("%s%s", logPrefix, message)
}

type string_pair struct {
	fi string
	se string
}

func makeRequest(method string, s string, headers []string_pair, body []byte) (bodyr []byte, resheader http.Header, StatusCode int, ok bool) {
	ok = true
	req, err := http.NewRequest(method, s, bytes.NewBuffer(body))
	if err != nil {
		log(err.Error())
		ok = false
		return
	}

	for _, header := range headers {
		req.Header.Set(header.fi, header.se)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log(err.Error())
		ok = false
		return
	}
	defer resp.Body.Close()
	StatusCode = resp.StatusCode
	resheader = resp.Header

	bodyr, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log(err.Error())
		ok = false
	}

	return
}

type TimeProvider struct {
	mu   sync.Mutex
	time time.Time
}

func (tp *TimeProvider) Stored() time.Time {
	tp.mu.Lock()
	defer tp.mu.Unlock()
	return tp.time
}

func (tp *TimeProvider) Set(t time.Time) {
	tp.mu.Lock()
	defer tp.mu.Unlock()
	tp.time = t
}

var tlock TimeProvider

const waitTime = 70

func (l *DefaultLogger) sendMessage(message string) (alterm string, request_code int) {
	// returns alterm, good_request
	const MAX_MESSAGE_SIZE = 2000
	var messageToSend string
	stored := tlock.Stored()
	elapsed := time.Since(stored)
	if elapsed < time.Duration(waitTime)*time.Millisecond {
		alterm = message
		request_code = 418
		return
	}

	headers := []string_pair{{fi: "Content-Type", se: "application/json"}}
	if len(message) > MAX_MESSAGE_SIZE {
		alterm = message[MAX_MESSAGE_SIZE:]
		messageToSend = message[:MAX_MESSAGE_SIZE]
	} else {
		alterm = ""
		messageToSend = message[:]
	}
	logf("Current msg size: %d\n", len(alterm))
	newmessage := "{\"content\":\"" + messageToSend + "\",\"username\":\"" + l.name + "\"}"
	_, _, StatusCode, ok := makeRequest("POST", l.url, headers, []byte(newmessage))
	if !ok {
		log("error when making request")
		alterm = message
		request_code = 400
		return
	}

	if StatusCode >= 300 {
		logf(messageToSend+"\n%d %v\n", StatusCode, ok)
		alterm = message
	}

	tlock.Set(time.Now())
	logf(messageToSend+"\n%d %v\n", StatusCode, ok)
	request_code = StatusCode
	return
}
