package kinspiration

import "os"

type Config struct {
	FilePath string
}

func (c *Config) Init() {
	quotesFile := os.Getenv("QUOTES_FILE")
	c.FilePath = quotesFile
}
