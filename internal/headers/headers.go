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
			break
		}

		re := regexp.MustCompile(`^\w`)
		if !re.Match(data) {
			return 0, false, fmt.Errorf("headers malformed")
		}

		parts := strings.Fields(string(line))

		validFieldName := regexp.MustCompile("^[a-zA-Z0-9!#\\$%&'\\*\\+\\-\\.\\^_`\\|,~]+:$")
		if !validFieldName.Match([]byte(parts[0])) {
			return 0, false, fmt.Errorf("headers malformed")
		}

		headerName, _ := strings.CutSuffix(parts[0], ":")
		headerValue := parts[1]

		h.Set(headerName, headerValue)

	}

	bytesConsumed := len(data) - 2
	return bytesConsumed, false, nil

}

func (h Headers) Get(nonNormalizedKey string) string {
	key := h.normalizeKey(nonNormalizedKey)
	if _, exists := h[key]; exists {
		return h[key]
	}
	return ""
}
func (h Headers) Set(nonNormalizedKey string, value string) {
	key := h.normalizeKey(nonNormalizedKey)
	if _, exists := h[key]; exists {
		h[key] = h[key] + " " + value
		return
	}
	h[key] = value
}

func (h Headers) normalizeKey(key string) string {
	return strings.ToLower(key)
}

func (h Headers) Update(key string, value string) {
	h[key] = value
}
