package telegram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateConfig(t *testing.T) {
	c := NewConfigWithToken("abcxyz")
	assert.Equal(t, c.Token, "abcxyz")
	assert.Equal(t, c.Endpoint(), "https://api.telegram.org/botabcxyz")
}
