package gomail

import "github.com/socious-io/gomq"

// The global config variable holds the configuration for the application.
var config = new(Config)

// Config represents the configuration structure for the Sendgrid Mail.
type Config struct {
	ApiKey         string
	Url            string
	Disabled       bool
	DefaultFrom    string
	DefaultSubject string
	Templates      map[string]string
	WorkerChannel  string
	MessageQueue   gomq.MessageQueue
}

// Setup initializes the Mail SDK with the provided configuration.
func Setup(cfg Config) error {
	// Set the global configuration to the provided config.
	config = &cfg
	return nil // Return nil to indicate successful setup.
}

func GetConfig() *Config {
	return config
}
