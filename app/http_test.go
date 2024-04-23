package main

import (
	"fmt"
	"testing"
)

func TestParseHttpRequest(t *testing.T) {
	test := struct {
		request string
		expect  HttpRequest
	}{
		request: "GET /echo/abc HTTP/1.1\r\n" +
			"Host: localhost:4221\r\n" +
			"User-Agent: curl/7.64.1\r\n" +
			"\n\r\n",
		expect: HttpRequest{
			Method:  "GET",
			Target:  "/echo/abc",
			Version: "HTTP/1.1",
			Headers: map[string]string{
				"Host":       "localhost:4221",
				"User-Agent": "curl/7.64.1",
			},
		},
	}

	got := ParseHttpRequest(test.request)

	expect := test.expect

	if got.Method != expect.Method {
		fmt.Printf("Expected Method: %v; Got: %v\n", expect.Method, got.Method)
	}

	if got.Target != expect.Target {
		fmt.Printf("Expected Target: %v; Got: %v\n", expect.Target, got.Target)
	}

	if got.Version != expect.Version {
		fmt.Printf("Expected Version: %v; Got: %v\n", expect.Version, got.Version)
	}

	for k, v := range expect.Headers {
		if got.Headers[k] != v {
			fmt.Printf("Expected: %v; Got: %v\n", v, got.Headers[k])
		}
	}
}
