package db

import (
	"context"
	"sync"

	"github.com/nmramorov/gophkeeper/internal/api/models"
)

type InMemoryDB struct {
	DBAPI
	Users       sync.Map
	Credentials sync.Map
}

func (d *InMemoryDB) SaveCredentials(ctx context.Context, data models.CredentialsData) error {
	select {
	case <-ctx.Done():
		return ErrContextTimeout
	default:
		d.Credentials.Store(data.UUID, data)
		return nil
	}
}

func (d *InMemoryDB) LoadCredentials(ctx context.Context, id string) (models.CredentialsData, error) {
	var result models.CredentialsData
	select {
	case <-ctx.Done():
		return result, ErrContextTimeout
	default:
		data, ok := d.Credentials.Load(id)
		if !ok {
			return result, ErrInMemoryDB
		}
		result = data.(models.CredentialsData)
		return result, nil
	}
}

func (d *InMemoryDB) Ping(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ErrDatabaseUnreachable
	default:
		return nil
	}
}

func (d *InMemoryDB) SaveUser(ctx context.Context, data models.User) error {
	select {
	case <-ctx.Done():
		return ErrContextTimeout
	default:
		d.Users.Store(data.Token, data)
		return nil
	}
}

func (d *InMemoryDB) LoadUser(ctx context.Context, token string) (models.User, error) {
	select {
	case <-ctx.Done():
		return models.User{}, ErrContextTimeout
	default:
		data, ok := d.Users.Load(token)
		if !ok {
			return models.User{}, ErrUserNotFound
		}
		return data.(models.User), nil
	}
}
