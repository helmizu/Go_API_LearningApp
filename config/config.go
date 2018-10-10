package config

//Config DB
type Config struct {
	Server   string
	Database string
}

//Init DB
func (c *Config) Init() {
	c.Server = "localhost"
	c.Database = "testing"
}

func (c Config) Read() Config {
	return c
}
