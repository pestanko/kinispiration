package kinspiration

import "os"

type Config struct {
	FilePath string
}

func (c *Config) Init() {
	c.FilePath = os.Getenv("QUOTES_FILE")
}
