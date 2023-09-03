package db

import (
	"context"

	"github.com/nmramorov/gophkeeper/internal/api/models"
)

type DBAPI interface {
	SaveCredentials(ctx context.Context, data models.CredentialsData) error
	LoadCredentials(ctx context.Context, id string) (models.CredentialsData, error)
	SaveUser(ctx context.Context, data models.User) error
	LoadUser(ctx context.Context, login string) (models.User, error)
	Ping(ctx context.Context) error
}
