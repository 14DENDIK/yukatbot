package config

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
	WebhookURL  string `toml:"webhook_url"`
	Token       string `toml:"token"`
}

// New ...
func New() *Config {
	return &Config{
		BindAddr:    ":3000",
		DatabaseURL: "host=localhost port=5432 user=sardor password=1 dbname=yukatdb",
		WebhookURL:  "",
		Token:       "",
	}
}
