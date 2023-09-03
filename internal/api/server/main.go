package server

import (
	"context"
	"fmt"

	"github.com/nmramorov/gophkeeper/internal/api/models"
	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
)

type StorageDB struct {
	db.DBAPI
}

type StorageServer struct {
	pb.UnimplementedStorageServer
	Storage StorageDB
}

func (s *StorageServer) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	var response pb.LoginUserResponse

	_, err := s.Storage.LoadUser(ctx, in.User.Login)
	if err != nil {
		response.Error = fmt.Sprintf("User with login %s already exists", in.User.Login)
	} else {
		response.Token = "token"
	}
	return &response, nil
}

func (s *StorageServer) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	var response pb.RegisterUserResponse
	newUser := models.User{
		Login: in.User.Login,
		Password: in.User.Password,
	}
	err := s.Storage.SaveUser(ctx, newUser)
	if err != nil {
		response.Error = fmt.Sprintf("Registration error with %s", in.User.Login)
	}
	return &response, nil
}
