package request

import (
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

type Request struct {
	RequestLine RequestLine
}

var ERR_MALFORMED_HTTP_REQUEST = fmt.Errorf("malformed http request")

func parseHttpRequest(b string) (*RequestLine, error) {
	idx := strings.Index(b, "\r\n")

	if idx == -1 {
		return nil, nil
	}

	firstLine := b[:idx]

	parts := strings.Split(firstLine, " ")

	if len(parts) != 3 {
		return nil, ERR_MALFORMED_HTTP_REQUEST
	}

	return &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   parts[2],
	}, nil
}

func RequestLineReader(reader io.Reader) (*Request, error) {
	b, err := io.ReadAll(reader)

	if err != nil {
		return nil, fmt.Errorf("Can't read all lines")
	}

	rl, err := parseHttpRequest(string(b))

	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *rl,
	}, nil
}
