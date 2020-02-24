package apiserver

// Configuration for API server
type Config struct {
	BindAddr string `tom;:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

//Function return created configuration
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}