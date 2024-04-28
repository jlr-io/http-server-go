package main

import (
	"fmt"
	"strings"
)

type Headers = map[string]string

const (
	Host          = "Host"
	UserAgent     = "User-Agent"
	ContentType   = "Content-Type"
	ContentLength = "Content-Length"
)

const (
	PlainText = "text/plain"
)

type HttpRequest struct {
	Method  string
	Target  string
	Version string
	Headers Headers
	Body    string
}

func ParseHttpRequest(message string) HttpRequest {
	request := HttpRequest{
		Headers: make(map[string]string),
	}
	lines := strings.Split(message, "\r\n")
	request.parseStart(lines[0])
	for _, l := range lines[1 : len(lines)-1] {
		if l == "" || l == "\r\n" {
			break
		}
		request.parseHeader(l)
	}
	request.Body = lines[len(lines)-1]
	return request
}
func (r *HttpRequest) parseStart(start string) {
	segments := strings.Split(start, " ")
	r.Method = segments[0]
	r.Target = segments[1]
	r.Version = segments[2]
}

func (r *HttpRequest) parseHeader(line string) {
	split := strings.SplitN(line, ":", 2)
	r.Headers[strings.ToLower(split[0])] = strings.TrimSpace(split[1])
}

func (r *HttpRequest) parseHeaders(headerLines []string) {
	headers := make(Headers)
	for _, h := range headerLines {
		split := strings.SplitN(h, ":", 2)
		headers[split[0]] = strings.TrimSpace(split[1])
	}
	r.Headers = headers
}

type HttpResponse struct {
	Version    string
	StatusCode string
	StatusText string
	Headers    Headers
	Body       string
}

const (
	Http1 = "HTTP/1.1"
)

func new200Response() HttpResponse {
	return HttpResponse{
		Version:    Http1,
		StatusCode: "200",
		StatusText: "OK",
		Headers:    make(Headers),
	}
}

func new201Response() HttpResponse {
	return HttpResponse{
		Version:    Http1,
		StatusCode: "201",
		StatusText: "CREATED",
		Headers:    make(Headers),
	}
}

func new404Response() HttpResponse {
	return HttpResponse{
		Version:    Http1,
		StatusCode: "404",
		StatusText: "NOT FOUND",
		Headers:    make(Headers),
	}
}

func (h HttpResponse) Encode() []byte {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s %s %s\r\n", h.Version, h.StatusCode, h.StatusText))

	for k, v := range h.Headers {
		sb.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}

	sb.WriteString("\r\n")

	sb.WriteString(fmt.Sprintf("%s\r\n", h.Body))

	return []byte(sb.String())
}
