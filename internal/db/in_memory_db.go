package db

import (
	"context"
	"sync"

	"github.com/nmramorov/gophkeeper/internal/api/models"
)

type InMemoryDB struct {
	DBAPI
	users       sync.Map
	credentials sync.Map
}

func (d *InMemoryDB) SaveCredentials(ctx context.Context, data models.CredentialsData) error {
	select {
	case <-ctx.Done():
		return ErrContextTimeout
	default:
		d.credentials.Store(data.UUID, data)
		return nil
	}
}

func (d *InMemoryDB) LoadCredentials(ctx context.Context, id string) (models.CredentialsData, error) {
	select {
	case <-ctx.Done():
		return models.CredentialsData{}, ErrContextTimeout
	default:
		data, ok := d.credentials.Load(id)
		if !ok {
			return models.CredentialsData{}, ErrInMemoryDB
		}
		return data.(models.CredentialsData), nil
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
		d.users.Store(data.Login, data)
		return nil
	}
}

func (d *InMemoryDB) LoadUser(ctx context.Context, login string) (models.User, error) {
	select {
	case <-ctx.Done():
		return models.User{}, ErrContextTimeout
	default:
		data, ok := d.users.Load(login)
		if !ok {
			return models.User{}, ErrInMemoryDB
		}
		return data.(models.User), nil
	}
}
