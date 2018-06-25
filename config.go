package main

import (
	"os"
)

// Config Represents database server and credentials
type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
	c.Server = os.Getenv("SERVER_NAME")
	c.Database = os.Getenv("DATABASE_NAME")
}
