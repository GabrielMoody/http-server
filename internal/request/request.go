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

	if len(parts) != 3 || !isHttpMethodValid(parts[0]) || !isHttpVersionValid(parts[2]) {
		return nil, ERR_MALFORMED_HTTP_REQUEST
	}

	return &RequestLine{
		Method:        parts[0],
		RequestTarget: parts[1],
		HttpVersion:   parts[2],
	}, nil
}

func isHttpVersionValid(m string) bool {
	if strings.ToUpper(m) == "HTTP/1.1" {
		return true
	}

	return false
}

func isHttpMethodValid(m string) bool {
	m = strings.ToUpper(m)

	if m == "GET" || m == "POST" || m == "PUT" || m == "PATCH" || m == "DELETE" || m == "HEAD" || m == "CONNECT" || m == "OPTIONS" || m == "TRACE" {
		return true
	}

	return false
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
