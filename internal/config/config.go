package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	RabbitMQ struct {
		URI string `json:"uri"`
	} `json:"rabbitmq"`
}

var AppConfig Config

func LoadConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	fmt.Println("Environment: ", env)

	// Open the config file
	fileName := "config." + env + ".json"
	configFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open config file: %v", err)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("could not decode config JSON: %v", err)
	}

	log.Println("Configuration loaded successfully")
}
