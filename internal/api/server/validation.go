package server

import (
	"context"

	"github.com/nmramorov/gophkeeper/internal/db"
)

func ValidateToken(incoming, stored string) bool {
	return incoming == stored
}

func (s *StorageServer) ValidateRequest(ctx context.Context, token string) string {
	mctx, mcancel := mergeContext(ctx, s.gctx)

	defer mcancel()

	decodedToken := DecodeToken(token)
	user, err := s.Storage.FindUser(mctx, decodedToken.Login, decodedToken.Password)
	switch err {
	case db.ErrUserNotFound:
		return "authorization error: wrong username"
	case db.ErrContextTimeout:
		return "context timeout. internal server error"
	case db.ErrInMemoryDB:
		return "internal server error"
	}
	for _, session := range user.Sessions {
		if ValidateToken(token, session.Token) {
			// In case of successfull validation value for response.Error is empty
			return ""
		}
	}

	// No tokens registered for the user, therefore token is invalid
	return "authorization error: invalid token"
}
