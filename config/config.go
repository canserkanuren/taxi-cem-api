package config

import (
	"os"
)

var Config *Configuration

type Configuration struct {
	Server *ServerConfiguration
}

type ServerConfiguration struct {
	Port     string
	SendMail *SendMailConfiguration
	CorsUrl  string
}

type SendMailConfiguration struct {
	Key         string
	DefaultMail string
}

func Setup(configPath string) {
	var configuration *Configuration
	// viper.SetConfigFile(configPath)
	// viper.SetConfigType("yaml")

	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file, %s", err)
	// }

	// err := viper.Unmarshal(&configuration)

	// if err != nil {
	// 	log.Fatalf("Unable to decode into struct, %v", err)
	// }

	port := os.Getenv("PORT")

	configuration = new(Configuration)
	configuration.Server = new(ServerConfiguration)
	configuration.Server.Port = port
	configuration.Server.SendMail = new(SendMailConfiguration)
	configuration.Server.SendMail.Key = os.Getenv("SENDMAIL_KEY")
	configuration.Server.SendMail.DefaultMail = os.Getenv("SENDMAIL_DEFAULTMAIL")
	configuration.Server.CorsUrl = os.Getenv("CORS_URL")
	Config = configuration
}

func GetConfig() *Configuration {
	return Config
}
