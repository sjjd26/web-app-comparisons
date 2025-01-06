package database

import (
	"context"
	"web-app/auth-api/models"
)

// rename package to database and type to Context -> database.Context
type DbContext interface {
	Close() error
	InsertUser(ctx context.Context, user models.User) (int, error)
}
