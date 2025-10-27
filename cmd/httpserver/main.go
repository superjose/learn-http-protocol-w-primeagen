package main

import (
	"GO_HTTP_PROTOCOL/internal/request"
	"GO_HTTP_PROTOCOL/internal/response"
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
	server, err := server.Serve(port, func(res *response.Response, req *request.Request) *server.HandleError {
		res.Headers.Update("Content-Type", "text/html")
		if req.RequestLine.RequestTarget == "/yourproblem" {
			f, _ := os.ReadFile("./cmd/httpserver/static/400.html")
			res.Status = response.HTTP_400
			res.Body.Write(f)
			return nil
		}
		if req.RequestLine.RequestTarget == "/myproblem" {
			f, _ := os.ReadFile("./cmd/httpserver/static/500.html")
			res.Status = response.HTTP_500
			res.Body.Write(f)
			return nil
		}
		f, _ := os.ReadFile("./cmd/httpserver/static/200.html")
		res.Status = response.HTTP_200
		res.Body.Write(f)
		return nil
	})
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
