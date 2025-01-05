package databasecontext

import (
	"context"
	"errors"
	"log"
	models "web-app/auth-api/Models"

	"github.com/jackc/pgx/v5"
)

type PgxDatabaseContext struct {
	dbConfig *pgx.ConnConfig
	db       *pgx.Conn
}

func NewPgxDatabaseContext(connectionString string) *PgxDatabaseContext {
	dbConfig, err := pgx.ParseConfig(connectionString)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	db, err := pgx.ConnectConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}

	return &PgxDatabaseContext{
		db:       db,
		dbConfig: dbConfig,
	}
}

func (db *PgxDatabaseContext) Close() error {
	if db.db != nil {
		db.db.Close(context.Background())
		return nil
	} else {
		return errors.New("database connection is not open")
	}
}

func (db *PgxDatabaseContext) InsertUser(ctx context.Context, user models.User) (int, error) {
	// TODO: Implement this method
	log.Println("CreateUser method not implemented")
	return 0, nil
}
