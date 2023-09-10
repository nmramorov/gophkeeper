package server

import (
	"context"

	"github.com/nmramorov/gophkeeper/internal/db"
)

func ValidateToken(incoming, stored string) bool {
	return incoming == stored
}

func (s *StorageServer) ValidateRequest(ctx context.Context, token string) string {
	decodedToken := DecodeToken(token)
	user, err := s.Storage.FindUser(ctx, decodedToken.Login, decodedToken.Password)
	switch err {
	case db.ErrUserNotFound:
		return "authorization error: wrong username"
	case db.ErrContextTimeout:
		return "context timeout. internal server error"
	case db.ErrInMemoryDB:
		return "internal server error"
	}
	if !ValidateToken(token, user.Token) {
		return "authorization error: invalid token"
	}
	// In case of successfull validation value for response.Error is empty
	return ""
}
