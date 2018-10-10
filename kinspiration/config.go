package kinspiration

import "os"

type Config struct {
	QuotesPath string
}

func (c *Config) Init() {
	quotesFile := os.Getenv("QUOTES_FILE")
	if quotesFile == "" {
		quotesFile = "quotes.json"
	}
	c.QuotesPath = quotesFile
}
