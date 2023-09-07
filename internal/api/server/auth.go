package server

import (
	"context"
	"fmt"

	"github.com/nmramorov/gophkeeper/internal/api/models"
	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
)

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

func (s *StorageServer) Register(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	var response pb.RegisterUserResponse
	token := GenerateToken(in.User.Login, in.User.Password)
	newUser := models.User{
		Login:    in.User.Login,
		Password: in.User.Password,
		Token:    token,
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
