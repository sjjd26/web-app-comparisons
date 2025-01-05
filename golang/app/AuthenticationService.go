package main

import (
	"fmt"
	"log"
	"net/http"
	databasecontext "web-app/auth-api/DatabaseContext"
	handlers "web-app/auth-api/Handlers"
)

type AuthenticationService struct {
	databaseContext databasecontext.DatabaseContext
}

func (service AuthenticationService) RegisterHandlers() {
	// Users
	userHandler := handlers.NewUserHandler(service.databaseContext)
	http.HandleFunc("POST /register", userHandler.RegisterNewUser)
}

func (service AuthenticationService) StartServer(port string) {
	fmt.Println("Starting server. Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
