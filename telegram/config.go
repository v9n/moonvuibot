package telegram

import (
	"fmt"
	"strings"
)

type Config struct {
	Token string
	url   string
}

func NewConfig(token string, options ...func(*Config)) *Config {
	c := Config{token, "https://api.telegram.org/bot"}
	for _, option := range options {
		option(&c)
	}
	return &c
}

func (c *Config) WithUrl(value string) {
	c.url = value
	if !strings.HasSuffix(c.url, "/") {
		c.url = c.url + "/"
	}
}

func (c *Config) Endpoint(method string) string {
	if method == "" {
		return fmt.Sprintf("%s%s", c.url, c.Token)
	}

	return fmt.Sprintf("%s%s/%s", c.url, c.Token, method)
}
