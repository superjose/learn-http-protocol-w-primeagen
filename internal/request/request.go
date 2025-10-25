package request

import (
	"GO_HTTP_PROTOCOL/internal/headers"
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
	Headers     headers.Headers
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	read := bufio.NewReader(reader)
	i := 0
	headers := headers.NewHeaders()
	req := Request{
		Headers: headers,
	}

	for {
		line, _, err := read.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Printf("%s\n", err)
			break
		}

		if len(line) == 0 {
			break
		}

		if i == 0 {
			s := string(line)
			arrs := strings.Fields(s)

			if len(arrs) != 3 {
				return nil, fmt.Errorf("Request malformed")
			}

			httpVersion, _ := strings.CutPrefix(arrs[2], "HTTP/")

			req.RequestLine = RequestLine{
				HttpVersion:   httpVersion,
				RequestTarget: arrs[1],
				Method:        arrs[0],
			}
		} else {
			_, _, err := headers.Parse(line)
			if err != nil {
				fmt.Printf("The error - %s", err)
			}
		}
		i++
	}
	return &req, nil
}
