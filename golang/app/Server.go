package main

import (
	"fmt"
	"os"
	databasecontext "web-app/auth-api/DatabaseContext"
)

const PORT_NUMBER = "8000"

func main() {
	// Connect to the database
	connectionString := os.Getenv("DATABASE_URL")
	fmt.Println("Connecting to pgbouncer: " + connectionString)

	databaseContext := databasecontext.NewPgxDatabaseContext(connectionString)
	defer databaseContext.Close()

	// Start the server
	authenticationService := AuthenticationService{databaseContext: databaseContext}
	authenticationService.RegisterHandlers()
	authenticationService.StartServer(PORT_NUMBER)
}
