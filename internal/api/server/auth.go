package server

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nmramorov/gophkeeper/internal/api/models"
	"github.com/nmramorov/gophkeeper/internal/db"
	pb "github.com/nmramorov/gophkeeper/internal/proto"
)

func (s *StorageServer) Login(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	mctx, mcancel := mergeContext(ctx, s.gctx)

	defer mcancel()

	var response pb.LoginUserResponse

	token := Token{
		Login:    in.User.Login,
		Password: in.User.Password,
		salt:     "salt",
	}
	encodedToken := EncodeToken(token)
	_, err := s.Storage.FindUser(mctx, in.User.Login, in.User.Password)
	switch err {
	case nil:
		response.Token = encodedToken
	case db.ErrUserNotFound:
		response.Error = fmt.Sprintf("user %s not found", in.User.Login)
	default:
		response.Error = fmt.Sprintf("error loading user %s: %e", in.User.Login, err)
	}
	return &response, nil
}

func (s *StorageServer) Register(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	mctx, mcancel := mergeContext(ctx, s.gctx)

	defer mcancel()

	var response pb.RegisterUserResponse

	token := Token{
		Login:    in.User.Login,
		Password: in.User.Password,
		salt:     "salt",
	}
	encodedToken := EncodeToken(token)
	user, err := s.Storage.FindUser(mctx, in.User.Login, in.User.Password)
	switch err {
	case nil:
		response.Error = fmt.Sprintf("user already exists %s", in.User.Login)
	case db.ErrContextTimeout:
		response.Error = fmt.Sprintf("registration. find user timeout %s: %e", in.User.Login, err)
	case db.ErrInMemoryDB:
		response.Error = fmt.Sprintf("internal server error %s: %e", in.User.Login, err)
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
	}
	return &response, nil
}
