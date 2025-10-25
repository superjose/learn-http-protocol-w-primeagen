package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const port = ":42069"

func readBuffer(reader io.Reader) {
	scanner := bufio.NewReader(reader)

	for {
		// buffer := make([]byte, 8)
		line, _, errBytes := scanner.ReadLine()
		if errors.Is(errBytes, io.EOF) {
			return
		} else if errBytes != nil {
			fmt.Printf("Errored: %s", errBytes)
			return
		}
		fmt.Printf("%s\n", line)
	}
}

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
			readBuffer(c)
			c.Close()
		}(conn)

	}
	// readBuffer(file)

}
