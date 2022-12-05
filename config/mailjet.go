package config

import (
	"os"
)

type MailjetConfig struct {
	APIKey      string
	SecretKey   string
	SenderName  string
	SenderEmail string
}

func LoadMailjetConfig() MailjetConfig {
	return MailjetConfig{
		APIKey:      os.Getenv("MAILJET_API_KEY"),
		SecretKey:   os.Getenv("MAILJET_SECRET_KEY"),
		SenderName:  os.Getenv("MAILJET_SENDER_NAME"),
		SenderEmail: os.Getenv("MAILJET_SENDER_EMAIL"),
	}
}
