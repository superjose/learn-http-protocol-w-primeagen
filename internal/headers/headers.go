package headers

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Headers map[string]string

func NewHeaders() Headers {
	return make(Headers)
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	stream := bufio.NewReader(bytes.NewReader(data))

	for {
		streamLine, _, err := stream.ReadLine()

		line := bytes.Trim(streamLine, " ")

		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return 0, false, err
		}
		if len(line) == 0 {
			continue
		}

		re := regexp.MustCompile(`^\w`)
		if !re.Match(data) {
			return 0, false, fmt.Errorf("headers malformed")
		}

		fmt.Println("🐥")
		fmt.Printf("%s\n", line)
		fmt.Println("🐥")

		parts := strings.Fields(string(line))

		headerName, _ := strings.CutSuffix(parts[0], ":")
		headerValue := parts[1]

		validFieldName := regexp.MustCompile("^[a-zA-Z0-9!#\\$%&'\\*\\+\\-\\.\\^_`\\|,~)]+$")
		if !validFieldName.Match([]byte(headerName)) {
			return 0, false, fmt.Errorf("headers malformed")
		}

		fmt.Printf("header name %s\n", headerName)
		fmt.Printf("header value %s\n", headerValue)

		if _, exists := h[headerName]; exists {
			h[headerName] = h[headerName] + " " + headerValue
			continue
		}
		h[headerName] = headerValue
	}

	bytesConsumed := len(data) - 2
	return bytesConsumed, false, nil

}
