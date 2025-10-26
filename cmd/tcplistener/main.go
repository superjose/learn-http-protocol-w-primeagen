package main

import (
	"GO_HTTP_PROTOCOL/internal/request"
	"fmt"
	"log"
	"net"
	"strings"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Error while opening the file")
		return
	}
	defer listener.Close()
	fmt.Printf("TCP Listening on port %s\n", port)
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			// readBuffer(c)
			req, err := request.RequestFromReader(c)
			if err != nil {
				log.Printf("[DEBUG] %v", err)
			}
			printRequest(req)
			c.Close()
		}(conn)

	}
	// readBuffer(file)

}

func printRequest(req *request.Request) {

	fmt.Println("Request Line:")
	fmt.Printf("- Method: %s\n", req.RequestLine.Method)
	fmt.Printf("- Target: %s\n", req.RequestLine.RequestTarget)
	fmt.Printf("- Version: %s\n", req.RequestLine.HttpVersion)

	fmt.Println("Headers:")

	for key, val := range req.Headers {
		fmt.Printf("- %s: %s\n", strings.ToUpper(key), strings.ToUpper(val))
	}

	fmt.Println("Body:")
	fmt.Printf("%s", req.Body)
}
