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

	_, err := s.Storage.LoadUser(ctx, in.User.Login)
	if err != nil {
		response.Error = fmt.Sprintf("Error loading user %s: %e", in.User.Login, err)
	} else {
		response.Token = "token"
	}
	return &response, nil
}

func (s *StorageServer) Register(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	var response pb.RegisterUserResponse
	newUser := models.User{
		Login:    in.User.Login,
		Password: in.User.Password,
	}
	user, err := s.Storage.LoadUser(ctx, in.User.Login)
	if err == db.ErrInMemoryDB {
		err = s.Storage.SaveUser(ctx, newUser)
		if err != nil {
			response.Error = fmt.Sprintf("Registration error with %s: %e", in.User.Login, err)
		}

		return &response, nil
	} else {
		response.Error = fmt.Sprintf("Internal server error: %e", err)
	}

	if user.Login == newUser.Login {
		response.Error = fmt.Sprintf("User %s already exists", in.User.Login)
	}

	return &response, nil
}
