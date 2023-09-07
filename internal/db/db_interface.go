package db

import (
	"context"

	"github.com/nmramorov/gophkeeper/internal/api/models"
)

type Credentials interface {
	SaveCredentials(ctx context.Context, data models.CredentialsData) error
	LoadCredentials(ctx context.Context, id string) (models.CredentialsData, error)
}

type Auth interface {
	SaveUser(ctx context.Context, data models.User) error
	LoadUser(ctx context.Context, token string) (models.User, error)
}

type DBAPI interface {
	Credentials
	Auth
	Ping(ctx context.Context) error
}
