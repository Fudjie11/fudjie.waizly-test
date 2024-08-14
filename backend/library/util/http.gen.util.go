package util

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type HttpRequestHeader struct {
	Key   string
	Value string
}

type HttpRequestData struct {
	Method        string
	Url           string
	JsonData      []byte
	CustomHeaders []HttpRequestHeader
}

func NewHttpRequest(httpRequest HttpRequestData) (*http.Response, error) {

	req, err := http.NewRequest(httpRequest.Method, httpRequest.Url, bytes.NewBuffer(httpRequest.JsonData))
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Error creating HTTP request:", err))
	}

	req.Header.Set("Content-Type", "application/json")
	if httpRequest.CustomHeaders != nil {
		for _, v := range httpRequest.CustomHeaders {
			req.Header.Set(v.Key, v.Value)

		}
	}

	client := &http.Client{
		Timeout: time.Second * 10, // Adjust timeout as needed
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Error sending request:", err))
	}

	defer resp.Body.Close()

	return resp, nil
}
