package service

import (
	"context"
)

type IEmailService interface {
	SendEmail(ctx context.Context, to, subject, body string) error
}
