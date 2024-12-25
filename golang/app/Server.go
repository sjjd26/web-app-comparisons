package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	AuthenticationHandlers "web-app/auth-api/Handlers"

	"github.com/jackc/pgx/v5"
)

const PORT = "8000"

func main() {
	// Connect to the database
	connectionString := os.Getenv("DATABASE_URL")
	fmt.Println("Connecting to pgbouncer: " + connectionString)

	dbConfig, err := pgx.ParseConfig(connectionString)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	db, err := pgx.ConnectConfig(context.Background(), dbConfig)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}

	defer func() {
		if db != nil {
			db.Close(context.Background())
		}
	}()

	// Test the connection
	fmt.Println("Testing connection")
	rows, err := db.Query(context.Background(), "SELECT * FROM flyway_schema_history")
	if err != nil {
		log.Fatalf("QueryRow.Values failed: %v\n", err)
	}
	defer rows.Close()

	// Get column metadata
	fieldDescriptions := rows.FieldDescriptions()
	columnCount := len(fieldDescriptions)
	fmt.Printf("Number of columns: %d\n", columnCount)

	// Process each row dynamically
	for rows.Next() {
		// Create a slice of empty interfaces to hold the column data
		values := make([]interface{}, columnCount)

		// Create references to the slice for Scan to populate
		valueRefs := make([]interface{}, columnCount)
		for i := range values {
			valueRefs[i] = &values[i]
		}

		// Scan the row into the slice
		if err := rows.Scan(valueRefs...); err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
		}

		// Print or process the row
		fmt.Println("Row:")
		for i, value := range values {
			fmt.Printf("  Column %s: %v\n", fieldDescriptions[i].Name, value)
		}
	}

	// Check for errors during iteration
	if rows.Err() != nil {
		log.Fatalf("Error iterating over rows: %v\n", rows.Err())
	}

	// --------------------------------------------------------------------------------

	// Register handlers
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, world!")
		io.WriteString(writer, "Hello, world!")
	})

	// Authentication
	http.HandleFunc("POST /register", AuthenticationHandlers.RegisterNewUser)

	// Start the server
	fmt.Println("Starting server. Listening on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
