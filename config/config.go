package config

import (
	"log"
	"os"
)

type Config struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
}

func LoadConfig() Config {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	callbackURL := os.Getenv("CALLBACK_URL")


	if clientID == "" || clientSecret == "" || callbackURL == "" {
		log.Fatalf("Missing required environment variables, %s%s%s\n", clientID, clientSecret, callbackURL)
	}

	return Config{
		ClientID: clientID,
		ClientSecret: clientSecret,
		CallbackURL: callbackURL,
	}
}
