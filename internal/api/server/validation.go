package server

import (
	"context"

	"github.com/nmramorov/gophkeeper/internal/db"
)

func ValidateToken(incoming, stored string) bool {
	return incoming == stored
}

func (s *StorageServer) ValidateRequest(ctx context.Context, token string) string {
	user, err := s.Storage.LoadUser(ctx, token)
	if err != nil {
		if err == db.ErrUserNotFound {
			return "authorization error: invalid token"
		} else {
			return "internal server error"
		}
	}
	if !ValidateToken(token, user.Token) {
		return "authorization error: invalid token"
	}
	// In case of successfull validation value for response.Error is empty
	return ""
}
