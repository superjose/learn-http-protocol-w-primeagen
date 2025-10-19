package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func readBuffer(reader io.Reader) {
	scanner := bufio.NewReader(reader)

	for {
		// buffer := make([]byte, 8)
		line, _, errBytes := scanner.ReadLine()
		if errors.Is(errBytes, io.EOF) {
			os.Exit(0)
		} else if errBytes != nil {
			fmt.Printf("Errored: %s", errBytes)
			os.Exit(1)
		}
		fmt.Printf("read: %s\n", line)
	}
}

func main() {
	file, err := os.OpenFile("./messages.txt", os.O_RDONLY, 0444)

	if err != nil {
		log.Fatal("Error while opening the file")
		return
	}
	defer file.Close()
	readBuffer(file)

}
