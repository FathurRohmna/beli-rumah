package config

import (
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func GetOAuthConfig() *oauth2.Config {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Fatalf("OAUTH_CLIENT_ID and OAUTH_CLIENT_SECRET must be set in environment variables")
	}

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{gmail.GmailSendScope},
		RedirectURL:  "http://localhost:8080",
	}
}
