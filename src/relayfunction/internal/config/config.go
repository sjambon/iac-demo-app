package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const defaultPort = "8080"

type Configuration struct {
	Port           string
	RelayTargetUrl string
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
		Port:           port,
		RelayTargetUrl: getEnvValue("RELAY_TARGET_URL"),
	}
}

func getEnvValue(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Default().Printf("Environment variable %s is not set", key)
	}

	return value
}
