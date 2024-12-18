package service

import (
	"beli-tanah/config"
	"beli-tanah/helper"
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailService struct {
	config *oauth2.Config
}

func NewEmailService() IEmailService {
	config := config.GetOAuthConfig()
	return &EmailService{config: config}
}

func (s *EmailService) SendEmail(ctx context.Context, to, subject, body string) error {
	token := config.GetToken()
	tokenSource := s.config.TokenSource(ctx, &token)
	service, err := gmail.NewService(ctx, option.WithTokenSource(tokenSource))
	if err != nil {
		return fmt.Errorf("failed to create Gmail service: %w", err)
	}

	message := fmt.Sprintf("From: me\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	msg := &gmail.Message{
		Raw: helper.EncodeBase64URLSafe([]byte(message)),
	}

	_, err = service.Users.Messages.Send("me", msg).Do()
	fmt.Print(err)
	return err
}
