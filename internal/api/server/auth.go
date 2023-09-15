package server

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nmramorov/gophkeeper/internal/api/models"
	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
)

func (s *StorageServer) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	var response pb.LoginUserResponse

	token := Token{
		Login:    in.User.Login,
		Password: in.User.Password,
		salt:     "salt",
	}
	encodedToken := EncodeToken(token)
	_, err := s.Storage.FindUser(ctx, in.User.Login, in.User.Password)
	switch err {
	case nil:
		response.Token = encodedToken
		return &response, nil
	case db.ErrUserNotFound:
		response.Error = fmt.Sprintf("user %s not found", in.User.Login)
		return &response, ErrUserNotFound
	default:
		response.Error = fmt.Sprintf("error loading user %s: %e", in.User.Login, err)
		return &response, ErrDatabaseError
	}
}

func (s *StorageServer) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	var response pb.RegisterUserResponse

	token := Token{
		Login:    in.User.Login,
		Password: in.User.Password,
		salt:     "salt",
	}
	encodedToken := EncodeToken(token)
	user, err := s.Storage.FindUser(ctx, in.User.Login, in.User.Password)
	switch err {
	case nil:
		response.Error = fmt.Sprintf("user already exists %s", in.User.Login)
		return &response, ErrUserAlreadyExists
	case db.ErrContextTimeout:
		response.Error = fmt.Sprintf("registration. find user timeout %s: %e", in.User.Login, err)
		return &response, ErrTimeout
	case db.ErrInMemoryDB:
		response.Error = fmt.Sprintf("internal server error %s: %e", in.User.Login, err)
		return &response, ErrDatabaseError
	}

	sessions := user.Sessions
	sessions = append(sessions, models.Session{
		UUID:  uuid.NewString(),
		Token: encodedToken,
	})

	newUser := models.User{
		UUID:     uuid.NewString(),
		Login:    in.User.Login,
		Password: in.User.Password,
		Sessions: sessions,
	}
	err = s.Storage.SaveUser(ctx, newUser)
	if err != nil {
		response.Error = fmt.Sprintf("registration error with %s: %e", in.User.Login, err)
		return &response, ErrDatabaseError
	}
	return &response, nil
}
