package headers

import (
	"fmt"
	"regexp"
	"strings"
)

type Headers map[string]string

func NewHeaders() Headers {
	return make(Headers)
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	line, _ := strings.CutSuffix(string(data), "\r\n\r\n")
	re := regexp.MustCompile(`^\w`)
	if !re.Match(data) {
		return 0, false, fmt.Errorf("headers malformed")
	}

	parts := strings.Fields(line)

	headerName, _ := strings.CutSuffix(parts[0], ":")
	headerValue := parts[1]

	fmt.Printf("header name %s\n", headerName)
	fmt.Printf("header value %s\n", headerValue)

	h[headerName] = headerValue

	bytesConsumed := len(data) - 2

	return bytesConsumed, false, nil

}
