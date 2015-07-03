package telegram

type Config struct {
	Token string
}

func NewConfigWithToken(token string) *Config {
	c := Config{token}
	return &c
}

func (c *Config) Endpoint(method string) string {
	return "https://api.telegram.org/bot" + c.Token + "/" + method
}
