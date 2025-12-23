package request

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseHttpRequestLine(t *testing.T) {
	rl, err := RequestLineReader(strings.NewReader("GET /test HTTP/1.1\r\n\r\n"))

	require.NoError(t, err)
	require.NotNil(t, rl)

	assert.Equal(t, "GET", rl.RequestLine.Method)
	assert.Equal(t, "/test", rl.RequestLine.RequestTarget)
	assert.Equal(t, "HTTP/1.1", rl.RequestLine.HttpVersion)

}
