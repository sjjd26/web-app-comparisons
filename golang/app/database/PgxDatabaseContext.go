package database

import (
	"context"
	"errors"
	"log"
	"web-app/auth-api/models"

	"github.com/jackc/pgx/v5"
)

type PgxContext struct {
	connConfig *pgx.ConnConfig
	conn       *pgx.Conn
}

func NewPgxContext(connectionString string) *PgxContext {
	connConfig, err := pgx.ParseConfig(connectionString)
	if err != nil {
		log.Fatalf("unable to parse DATABASE_URL: %v\n", err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		log.Fatalf("unable to connect to the database: %v\n", err)
	}

	return &PgxContext{
		conn:       conn,
		connConfig: connConfig,
	}
}

func (db *PgxContext) Close() error {
	if db.conn != nil {
		db.conn.Close(context.Background())
		return nil
	} else {
		return errors.New("database connection is not open")
	}
}

func (db *PgxContext) InsertUser(ctx context.Context, user models.User) (int, error) {
	query := `
		INSERT INTO users (email, password_hash, salt, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id
	`
	var userId int
	err := db.conn.QueryRow(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
		user.Salt,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&userId)

	if err != nil {
		return -1, err
	}

	return userId, nil
}
