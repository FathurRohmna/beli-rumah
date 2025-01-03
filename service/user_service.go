package service

import (
	"beli-tanah/model/web"
	"context"
)

type IUserService interface {
	Login(ctx context.Context, student web.LoginUserRequest) string
	Register(ctx context.Context, student web.RegisterUserRequest) web.UserResponse
	GetUserById(ctx context.Context, userId string) web.UserResponse
	GetMyDetail(ctx context.Context, userId string) web.MyDetailResponse
}
