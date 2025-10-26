package request

import (
	"GO_HTTP_PROTOCOL/internal/body"
	"GO_HTTP_PROTOCOL/internal/headers"
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Request struct {
	RequestLine RequestLine
	Headers     headers.Headers
	Body        []byte
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	read := bufio.NewReader(reader)
	req := Request{}

	// 1st - We extract the main line
	requestLine, err := extractRequest(read)

	if err != nil {
		return nil, err
	}

	req.RequestLine = requestLine

	requestHeaders, err := extractHeaders(read)

	if err != nil {
		return nil, err
	}

	req.Headers = requestHeaders

	contentLength := req.Headers.Get("Content-Length")
	if len(contentLength) == 0 {
		return &req, nil
	}

	requestBody, err := extractBody(read, contentLength)

	if err != nil {
		return nil, err
	}

	req.Body = requestBody

	return &req, nil
}

func extractRequest(reader *bufio.Reader) (RequestLine, error) {

	line, _, err := reader.ReadLine()

	// No error should come up, even EOF
	if err != nil {
		return RequestLine{}, err
	}

	s := string(line)
	arrs := strings.Fields(s)

	if len(arrs) != 3 {
		return RequestLine{}, fmt.Errorf("Request malformed")
	}

	httpVersion, _ := strings.CutPrefix(arrs[2], "HTTP/")

	requestLine := RequestLine{
		HttpVersion:   httpVersion,
		RequestTarget: arrs[1],
		Method:        arrs[0],
	}

	return requestLine, nil
}

func extractHeaders(reader *bufio.Reader) (headers.Headers, error) {
	headers := headers.NewHeaders()
	for {
		line, _, err := reader.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return headers, nil
		}

		if len(line) == 0 {
			break
		}

		_, _, err = headers.Parse(line)
		if err != nil {
			return headers, err
		}
	}
	return headers, nil
}

func extractBody(reader *bufio.Reader, contentLengthStr string) (body.Body, error) {
	body := body.NewBody()

	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return body, err
	}

	bodyBytes := make([]byte, contentLength)
	_, err = io.ReadFull(reader, bodyBytes)
	if err != nil {
		return nil, err
	}
	body.Parse(bodyBytes)
	return body, nil
}
