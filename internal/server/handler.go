package server

import (
	"GO_HTTP_PROTOCOL/internal/request"
	"GO_HTTP_PROTOCOL/internal/response"
	"fmt"
	"io"
	"log"
)

type HandleError struct {
	StatusCode response.StatusCode
	Message    string
}

type Handler func(res *response.Response, req *request.Request) *HandleError

func (he *HandleError) Write(w io.Writer) {
	a := func(code uint16, msg string) error {
		str := fmt.Sprintf("HTTP/1.1 %d %s\r\n", code, msg)
		log.Printf("[DEBUG] - %s\n", str)
		_, err := w.Write([]byte(str))
		return err
	}
	switch he.StatusCode {
	case response.HTTP_400:
		{
			a(400, "Bad Request")
		}
	case response.HTTP_500:
		{
			a(500, "Internal Server Error")
		}
	}
}
