package config

//Config DB
type Config struct {
	Server   string
	Database string
}

//Init DB
func (c *Config) Init() {
	c.Server = "localhost"
	c.Database = "GoLearn"
}

func (c Config) Read() Config {
	return c
}
