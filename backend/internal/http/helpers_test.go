package http

import (
	"bytes"
	"net/http"
)

func createTestRequest(method, target string, body []byte) *http.Request {
	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err != nil {
		return &http.Request{}
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req
}
