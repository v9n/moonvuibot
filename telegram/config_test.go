package telegram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateConfig(t *testing.T) {
	c := NewConfig("abcxyz")
	assert.Equal(t, c.Token, "abcxyz")
	assert.Equal(t, c.Endpoint(""), "https://api.telegram.org/botabcxyz")
	assert.Equal(t, c.Endpoint("method"), "https://api.telegram.org/botabcxyz/method")
}
