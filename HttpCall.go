package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpCall struct {
	Url string
}

func (hc *HttpCall) Call(requestBody []byte) (xml []byte, err error) {
	var (
		response *http.Response
		request  *http.Request
	)

	if request, err = http.NewRequest(http.MethodGet, hc.Url, bytes.NewBuffer(requestBody)); err != nil {
		return
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	if response, err = client.Do(request); err != nil {
		return
	}

	if xml, err = ioutil.ReadAll(response.Body); err != nil {
		return
	}

	return
}
