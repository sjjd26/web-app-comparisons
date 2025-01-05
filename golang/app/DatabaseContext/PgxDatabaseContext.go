package databasecontext

import (
	"context"
	"errors"
	"log"
	models "web-app/auth-api/Models"

	"github.com/jackc/pgx/v5"
)

type PgxDatabaseContext struct {
	connConfig *pgx.ConnConfig
	conn       *pgx.Conn
}

func NewPgxDatabaseContext(connectionString string) *PgxDatabaseContext {
	connConfig, err := pgx.ParseConfig(connectionString)
	if err != nil {
		log.Fatalf("unable to parse DATABASE_URL: %v\n", err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		log.Fatalf("unable to connect to the database: %v\n", err)
	}

	return &PgxDatabaseContext{
		conn:       conn,
		connConfig: connConfig,
	}
}

func (db *PgxDatabaseContext) Close() error {
	if db.conn != nil {
		db.conn.Close(context.Background())
		return nil
	} else {
		return errors.New("database connection is not open")
	}
}

func (db *PgxDatabaseContext) InsertUser(ctx context.Context, user models.User) (int, error) {
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
