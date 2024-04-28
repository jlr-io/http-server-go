package main

import "testing"

func TestEcho(t *testing.T) {
	test := struct {
		request HttpRequest
		expect  HttpResponse
	}{
		request: HttpRequest{
			Method:  "GET",
			Target:  "/echo/abc",
			Version: "HTTP/1.1",
			Headers: Headers{
				"Host":       "localhost:4221",
				"User-Agent": "curl/7.64.1",
			},
		},
		expect: HttpResponse{
			Version:    "HTTP/1.1",
			StatusCode: "200",
			StatusText: "OK",
			Headers: Headers{
				"Content-Type":   "text/plain",
				"Content-Length": "3",
			},
			Body: "abc",
		},
	}

	got := HandleEcho(test.request)

	expect := test.expect

	if got.Version != expect.Version {
		t.Errorf("Expected Version: %v; Got: %v\n", expect.Version, got.Version)
	}

	if got.StatusCode != expect.StatusCode {
		t.Errorf("Expected StatusCode: %v; Got: %v\n", expect.StatusCode, got.StatusCode)
	}

	if got.StatusText != expect.StatusText {
		t.Errorf("Expected StatusText: %v; Got: %v\n", expect.StatusText, got.StatusText)
	}

	for k, v := range expect.Headers {
		if got.Headers[k] != v {
			t.Errorf("Expected: %v; Got: %v\n", v, got.Headers[k])
		}
	}

	if got.Body != expect.Body {
		t.Errorf("Expected Body: %v; Got: %v\n", expect.Body, got.Body)
	}
}
