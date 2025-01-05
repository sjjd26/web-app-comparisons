package databasecontext

import (
	"context"
	models "web-app/auth-api/Models"
)

type DatabaseContext interface {
	Close() error
	InsertUser(ctx context.Context, user models.User) (int, error)
}
