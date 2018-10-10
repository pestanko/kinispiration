package kinspiration

import (
	"log"
	"os"
)

type Config struct {
	QuotesPath string
	AdminToken string
}

func (c *Config) setAdminToken() {
	adminToken := os.Getenv("QUOTES_TOKEN")
	c.AdminToken = adminToken
	if adminToken == "" {
		log.Printf("[CONFIG] Auth token has not been set - Admin API is not secured")
	} else {
		log.Printf("[CONFIG] Auth token has been set \"%s\" - Admin API is secured", adminToken)
	}
}

func (c *Config) setQuotesFile() {
	quotesFile := os.Getenv("QUOTES_FILE")

	if quotesFile == "" {
		quotesFile = "quotes.json"
		log.Printf("[CONFIG] Quotes database has not been set, using the default: %s", quotesFile)
	} else {
		log.Printf("[CONFIG] Quotes database has been set: %s", quotesFile)
	}

	c.QuotesPath = quotesFile
}

func (c *Config) Init() {
	c.setAdminToken()
	c.setQuotesFile()
}
