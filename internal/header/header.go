package header

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type Header map[string][]string

var ERR_MALFORMED_HEADERS = errors.New("malformed header request")

func ParseHeader(b string) (Header, error) {
	b = strings.TrimSpace(b)
	lines := strings.Split(b, "\r\n")

	header := make(Header)

	for _, val := range lines {
		if val == "" {
			continue
		}

		h := strings.SplitN(val, ":", 2)
		if len(h) != 2 {
			return nil, ERR_MALFORMED_HEADERS
		}

		key := strings.TrimSpace(h[0])
		values := strings.Split(h[1], ",")

		for i := range values {
			values[i] = strings.TrimSpace(values[i])
		}
		header[key] = append(header[key], values...)
	}

	return header, nil
}

func RequestHeaderReader(reader io.Reader) (Header, error) {
	b, err := io.ReadAll(reader)

	if err != nil {
		return nil, fmt.Errorf("Can't read all lines")
	}

	rl, err := ParseHeader(string(b))

	if err != nil {
		return nil, err
	}

	return rl, nil
}
