package config

import (
	"os"
	"time"

	"golang.org/x/oauth2"
)

func GetToken() oauth2.Token {
	accessToken := os.Getenv("ACCESS_TOKEN")
	refreshToken := os.Getenv("REFRESH_TOKEN")

	token := oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	return token
}
