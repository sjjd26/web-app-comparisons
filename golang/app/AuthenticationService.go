package main

import (
	"fmt"
	"log"
	"net/http"
	"web-app/auth-api/database"
	"web-app/auth-api/handlers"
)

type AuthenticationService struct {
	dbContext database.DbContext
}

func (service AuthenticationService) RegisterHandlers() {
	// Users
	userHandler := handlers.NewUserHandler(service.dbContext)
	http.HandleFunc("POST /register", userHandler.RegisterNewUser)
}

func (service AuthenticationService) StartServer(port string) {
	fmt.Println("Starting server. Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
