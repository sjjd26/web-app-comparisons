package databasecontext

import (
	"context"
	models "web-app/auth-api/Models"
)

type DatabaseContext interface {
	Close() error
	CreateUser(ctx context.Context, email string, password string) (models.User, error)
}
