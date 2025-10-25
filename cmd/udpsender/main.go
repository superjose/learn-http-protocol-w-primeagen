package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const address = "localhost:42069"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(">")
		userInput, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Input error! %s", err)
			continue
		}

		// fmt.Printf("Total length %d", len(userInput))

		_, errWrite := conn.Write([]byte(userInput))

		if errWrite != nil {
			fmt.Printf("Input error! %s", errWrite)
		}
		// fmt.Printf("Success %d", success)
	}

}
