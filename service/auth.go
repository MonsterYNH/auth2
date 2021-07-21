package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MonsterYNH/api/v1/oauth2"
	"github.com/MonsterYNH/athena/util"
	"github.com/MonsterYNH/auth2/database"
	"github.com/MonsterYNH/auth2/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthService struct {
	oauth2.UnimplementedAuth2SerivceServer
}

func (service *AuthService) Login(ctx context.Context, request *oauth2.Auth2LoginRequest) (*oauth2.Auth2LoginResponse, error) {
	conn := database.GetDatabase()

	user := new(models.User)
	if err := conn.Where("account=? and password=?", request.Account, request.Password).First(user).Error; err != nil {
		return nil, err
	}

	token, err := util.GenerateToken(user.ID, "jwt-key", time.Hour)
	if err != nil {
		return nil, err
	}

	if err := grpc.SendHeader(ctx, metadata.Pairs("athena_token", token)); err != nil {
		return nil, err
	}
	return &oauth2.Auth2LoginResponse{
		Account: user.Account,
		Name:    user.Name,
		Status:  true,
	}, nil
}

func (service *AuthService) Regist(ctx context.Context, request *oauth2.Auth2RegistRequest) (*oauth2.Auth2RegistResponse, error) {
	conn := database.GetDatabase()

	user := new(models.User)
	if err := conn.Model(user).Where("account=?", request.Account).Error; err != nil {
		return nil, err
	}

	if len(user.ID) != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("user account: %s is already exist", request.Account))
	}

	user.Account = request.Account
	user.Password = request.Password

	if err := conn.Create(user).Error; err != nil {
		return nil, err
	}

	return &oauth2.Auth2RegistResponse{
		Status: true,
	}, nil
}

func (service *AuthService) Auth(ctx context.Context, request *oauth2.Auth2AuthRequest) (*oauth2.Auth2AuthResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "no auth info")
	}

	values := md.Get("athena_token")
	if len(values) == 0 {
		return nil, status.Error(codes.Unauthenticated, "empty auth info")
	}
	userID, err := util.ParseToken(values[0], "jwt-key")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	conn := database.GetDatabase()
	user := new(models.User)
	if err := conn.Model(user).Where("id=?", userID).Error; err != nil {
		return nil, status.Error(codes.Unauthenticated, "wrong auth")
	}

	token, err := util.GenerateToken(user.ID, "jwt-key", time.Hour)
	if err != nil {
		return nil, err
	}

	if err := grpc.SendHeader(ctx, metadata.Pairs("athena_token", token)); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(err)
	return &oauth2.Auth2AuthResponse{
		Service: request.Service,
		Status:  true,
	}, nil
}
