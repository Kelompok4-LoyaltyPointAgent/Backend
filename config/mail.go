package config

import (
	"log"
	"os"
	"strconv"
)

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func LoadMailConfig() MailConfig {
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		log.Print("Invalid mail port")
	}

	return MailConfig{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     port,
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}
}
