package main

import (
	"GO_HTTP_PROTOCOL/internal/server"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const port = 42069

func main() {
	fmt.Printf("Listening on port %d\n", port)
	server, err := server.Serve(port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer server.Close()
	log.Println("Server started on port", port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}
