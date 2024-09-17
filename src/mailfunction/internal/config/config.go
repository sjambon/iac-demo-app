package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const defaultPort = "8080"

type Configuration struct {
	Port                     string
	MailServiceConfiguration mailServiceConfig
}

type mailServiceConfig struct {
	ApiKey     string
	SecretKey  string
	SenderMail string
}

func New() *Configuration {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err.Error())
		}
	}

	port, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !ok {
		port = defaultPort
	}

	return &Configuration{
		Port: port,
		MailServiceConfiguration: mailServiceConfig{
			ApiKey:     getEnvValue("MAILJET_API_KEY"),
			SecretKey:  getEnvValue("MAILJET_SECRET_KEY"),
			SenderMail: getEnvValue("MAILJET_SENDER_MAIL"),
		},
	}
}

func getEnvValue(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Default().Printf("Environment variable %s is not set", key)
	}

	return value
}
