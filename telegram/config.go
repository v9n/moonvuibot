package telegram

type Config struct {
	Token string
}

func NewConfigWithToken(token string) *Config {
	c := Config{token}
	return &c
}
