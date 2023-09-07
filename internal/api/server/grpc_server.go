package server

import (
	"context"
	"fmt"

	"github.com/nmramorov/gophkeeper/internal/api/models"
	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
)


type StorageServer struct {
	pb.UnimplementedStorageServer
	Storage db.DBAPI
}

func (s *StorageServer) Login(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	var response pb.LoginUserResponse

	token := GenerateToken(in.User.Login, in.User.Password)
	_, err := s.Storage.LoadUser(ctx, token)
	if err != nil {
		response.Error = fmt.Sprintf("error loading user %s: %e", in.User.Login, err)
	} else {
		response.Token = "token"
	}
	return &response, nil
}

func GenerateToken(login, password string) string {
	return login + password + "salt"
}

func (s *StorageServer) Register(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	var response pb.RegisterUserResponse
	token := GenerateToken(in.User.Login, in.User.Password)
	newUser := models.User{
		Login:    in.User.Login,
		Password: in.User.Password,
		Token: token,
	}
	user, err := s.Storage.LoadUser(ctx, token)
	if err == db.ErrUserNotFound {
		err = s.Storage.SaveUser(ctx, newUser)
		if err != nil {
			response.Error = fmt.Sprintf("registration error with %s: %e", in.User.Login, err)
		}

		return &response, nil
	} else {
		response.Error = fmt.Sprintf("internal server error: %e", err)
	}

	if user.Login == newUser.Login {
		response.Error = fmt.Sprintf("user %s already exists", in.User.Login)
	}

	return &response, nil
}

func ValidateToken(incoming, stored string) bool {
	return incoming == stored
}

func (s *StorageServer) SaveCredentials(ctx context.Context, in *pb.SaveCredentialsDataRequest) (*pb.SaveCredentialsDataResponse, error) {
	var response pb.SaveCredentialsDataResponse

	user, err := s.Storage.LoadUser(ctx, in.Token)
	if err != nil {
		if err == db.ErrUserNotFound {
			response.Error = "authorization error: invalid token"
			return &response, nil
		} else {
			response.Error = "internal server error"
			return &response, nil
		}
	}
	if !ValidateToken(in.Token, user.Token) {
		response.Error = "authorization error: invalid token"
		return &response, nil
	}

	newCredentials := models.CredentialsData{
		UUID: in.Data.Uuid,
		Login: in.Data.Login,
		Password: in.Data.Password,
	}
	err = s.Storage.SaveCredentials(ctx, newCredentials)
	if err != nil {
		response.Error = fmt.Sprintf("internal server error for data %s", in.Data.Uuid)
		return &response, nil
	}
	return &response, nil
}