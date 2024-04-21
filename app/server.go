package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
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

	request := string(buf[:n])
	lines := strings.Split(request, "\r\n")
	start := lines[0]
	path := strings.Split(start, " ")[1]

	response := "HTTP/1.1 404 NOT FOUND\r\n\r\n"
	if path == "/" {
		response = "HTTP/1.1 200 OK\r\n\r\n"
	}

	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Println("Failed to write to connection")
		os.Exit(1)
	}
}
