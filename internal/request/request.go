package request

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {

	read := bufio.NewReader(reader)
	i := 0
	req := Request{}
	for {
		cost, _, err := read.ReadLine()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		if i == 0 {
			s := string(cost)
			fmt.Printf("%s\n", s)
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
		}

		i++
	}
	return &req, nil
}
