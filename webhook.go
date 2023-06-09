package webhookLog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func log(s string) {
	fmt.Println(s)
}

func logf(s string, args ...interface{}) {
	fmt.Printf(s, args...)
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
	}

	for _, header := range headers {
		req.Header.Set(header.fi, header.se)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log(err.Error())
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

func (l *DefaultLogger) sendMessage(message string) (alterm string, request_code int) {
	// returns alterm, good_request
	const MAX_MESSAGE_SIZE = 2000
	var messageToSend string

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
	logf(messageToSend+"\n%d %v\n", StatusCode, ok)
	if !ok || StatusCode >= 300 {
		//bad request. do something?
	}
	request_code = StatusCode
	return
}
