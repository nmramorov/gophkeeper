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
	LoadUser(ctx context.Context, id string) (models.User, error)
}

type Binaries interface {
	SaveBinary(ctx context.Context, data models.BinaryData) error
	LoadBinary(ctx context.Context, id string) (models.BinaryData, error)
}

type Texts interface {
	SaveText(ctx context.Context, data models.TextData) error
	LoadText(ctx context.Context, id string) (models.TextData, error)
}

type Cards interface {
	SaveCard(ctx context.Context, data models.BankCardData) error
	LoadCard(ctx context.Context, id string) (models.BankCardData, error)
}

type Session interface {
	Add(ctx context.Context, session models.Session) error
	Remove(ctx context.Context, session models.Session) error
}

type DBAPI interface {
	Credentials
	Auth
	Texts
	Binaries
	Cards
	Ping(ctx context.Context) error
	FindUser(ctx context.Context, login, password string) (models.User, error)
}
