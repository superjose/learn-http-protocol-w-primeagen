package main

import (
	"GO_HTTP_PROTOCOL/internal/request"
	"GO_HTTP_PROTOCOL/internal/response"
	"GO_HTTP_PROTOCOL/internal/server"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

const port = 42069

func main() {
	fmt.Printf("Listening on port %d\n", port)
	server, err := server.Serve(port, func(res *response.Response, req *request.Request) *server.HandleError {
		res.Headers.Update("Content-Type", "text/html")
		beginsWith := regexp.MustCompile(`^/httpbin`)
		target := req.RequestLine.RequestTarget

		if target == "/yourproblem" {
			f, _ := os.ReadFile("./cmd/httpserver/static/400.html")
			res.Status = response.HTTP_400
			res.Body.Write(f)
			return nil
		}
		if target == "/myproblem" {
			f, _ := os.ReadFile("./cmd/httpserver/static/500.html")
			res.Status = response.HTTP_500
			res.Body.Write(f)
			return nil
		}

		/**
		Note that this doesn't support streaming
		*/
		if beginsWith.Match([]byte(target)) {
			route, _ := strings.CutPrefix(target, "/httpbin")
			res.Headers.Update("Transfer-Encoding", "chunked")
			fetch, err := http.Get("https://httpbin.org/" + route)
			if err != nil {
				hErr := &server.HandleError{
					StatusCode: response.HTTP_400,
					Message:    err.Error(),
				}
				return hErr
			}
			defer fetch.Body.Close()
			io.Copy(&res.Body, fetch.Body)
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
