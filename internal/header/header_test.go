package header

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseHeader(t *testing.T) {
	header := strings.NewReader("Host: localhost:8080\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n")

	parsedHeader, _ := RequestHeaderReader(header)

	assert.Equal(t, "localhost:8080", parsedHeader["Host"][0])
	assert.Equal(t, "curl/7.64.1", parsedHeader["User-Agent"][0])
	assert.Equal(t, "*/*", parsedHeader["Accept"][0])
}

func TestParseMalformedHeader(t *testing.T) {
	header := strings.NewReader("Host localhost:8080\r\nUser-Agent curl/7.64.1\r\nAccept */*\r\n\r\n")

	_, err := RequestHeaderReader(header)

	require.Error(t, err)
	assert.Equal(t, ERR_MALFORMED_HEADERS, err)
}

func TestParseHeaderMultipleValue(t *testing.T) {
	header := strings.NewReader("Accept: text/html\r\nAccept: application/json\r\n\r\n")

	parsedHeader, _ := RequestHeaderReader(header)

	assert.Equal(t, "text/html", parsedHeader["Accept"][0])
	assert.Equal(t, "application/json", parsedHeader["Accept"][1])
}

func TestParseHeaderInlineMultipleValue(t *testing.T) {
	header := strings.NewReader("Accept: text/html, application/json\r\n\r\n")

	parsedHeader, _ := RequestHeaderReader(header)

	assert.Equal(t, "text/html", parsedHeader["Accept"][0])
	assert.Equal(t, "application/json", parsedHeader["Accept"][1])
}
