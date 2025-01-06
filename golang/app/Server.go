package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"web-app/auth-api/database"
	"web-app/auth-api/middleware"
)

const PORT_NUMBER = "8000"

func main() {
	setupLogging()

	// Connect to the database
	connectionString := os.Getenv("DATABASE_URL")
	fmt.Println("Connecting to pgbouncer: " + connectionString)

	dbContext := database.NewPgxContext(connectionString)
	defer dbContext.Close()

	// Start the server
	authenticationService := AuthenticationService{dbContext: dbContext}
	authenticationService.RegisterHandlers()

	http.Handle("/", middleware.LoggingMiddleware(http.DefaultServeMux))
	authenticationService.StartServer(PORT_NUMBER)
}

func setupLogging() {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	var logOutput io.Writer

	switch os.Getenv("APP_ENV") {
	case "production":
		// Log errors to both stdout and file
		logOutput = io.MultiWriter(os.Stdout, logFile)
	default: // development
		// Log everything to stdout
		logOutput = os.Stdout
	}

	log.SetOutput(logOutput)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
