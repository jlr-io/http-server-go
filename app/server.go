package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Failed to close connection")
			os.Exit(1)
		}
	}(conn)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println("Failed to read from connection")
		os.Exit(1)
	}

	message := string(buf[:n])
	request := ParseHttpRequest(message)
	response := HttpResponse{}

	if strings.HasPrefix(request.Target, "/echo/") {
		response = HandleEcho(request)
	} else if strings.HasPrefix(request.Target, "/") {
		if request.Target == "/" {
			response = new200Response()
		} else {
			response = HandleHeader(request)
		}
	} else {
		response = new404Response()
	}

	_, err = conn.Write(response.Encode())
	if err != nil {
		log.Println("Failed to write to connection")
		os.Exit(1)
	}
}

func HandleEcho(request HttpRequest) HttpResponse {
	response := new200Response()
	response.Body = strings.TrimPrefix(request.Target, "/echo/")
	response.Headers = Headers{
		ContentType:   PlainText,
		ContentLength: fmt.Sprintf("%d", len(response.Body)),
	}
	return response
}

func HandleHeader(request HttpRequest) HttpResponse {
	header := strings.TrimPrefix(request.Target, "/")
	if headerValue, exists := request.Headers[header]; exists {
		response := new200Response()
		response.Body = headerValue
		response.Headers = Headers{
			ContentType:   PlainText,
			ContentLength: fmt.Sprintf("%d", len(response.Body)),
		}
		return response
	} else {
		return new404Response()
	}
}
