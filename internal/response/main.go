package response

import (
	"GO_HTTP_PROTOCOL/internal/headers"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
)

type StatusCode int

const (
	HTTP_200 StatusCode = iota
	HTTP_400
	HTTP_500
)

type Response struct {
	Status  StatusCode
	Body    bytes.Buffer
	Headers headers.Headers
}

func WriteStatusLine(w io.Writer, statusCode StatusCode) error {
	a := func(code uint16, msg string) error {
		str := fmt.Sprintf("HTTP/1.1 %d %s\r\n", code, msg)
		log.Printf("[DEBUG] - %s\n", str)
		_, err := w.Write([]byte(str))
		return err
	}
	switch statusCode {
	case HTTP_200:
		{
			return a(200, "OK")
		}
	case HTTP_400:
		{
			return a(400, "Bad Request")
		}
	case HTTP_500:
		{
			return a(500, "Internal Server Error")
		}
	}
	return nil
}

func GetDefaultHeaders(contentLen int) headers.Headers {
	headers := headers.NewHeaders()
	headers["Content-Length"] = strconv.Itoa(contentLen)
	headers["Connection"] = "close"
	headers["Content-Type"] = "text/plain"
	return headers
}

func WriteHeaders(w io.Writer, headers headers.Headers) error {
	for key, val := range headers {
		str := fmt.Sprintf("%s: %s\r\n", key, val)
		_, err := w.Write([]byte(str))
		if err != nil {
			return err
		}
	}
	_, err := w.Write([]byte("\r\n"))
	return err
}

func (res *Response) GetHeaders() headers.Headers {
	// This actually mutates it
	res.Headers.Update("Content-Length", strconv.Itoa(len(res.Body.Bytes())))
	hasTransferEncoding := len(res.Headers.Get("Transfer-Encoding")) > 0
	if hasTransferEncoding {
		res.Headers.Remove("Content-Length")
	}
	return res.Headers
}
func (res *Response) Write(w io.Writer) {
	w.Write(append(res.Body.Bytes(), "\n"...))
}
