package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"web-app/auth-api/database"
	"web-app/auth-api/models"
)

type UserHandler struct {
	dbContext database.DbContext
}

func NewUserHandler(dbContext database.DbContext) *UserHandler {
	return &UserHandler{
		dbContext: dbContext,
	}
}

func (handler UserHandler) RegisterNewUser(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")

	if email == "" || password == "" {
		http.Error(writer, "missing email or password", http.StatusBadRequest)
		return
	}

	newUser, err := models.NewUser(email, password)
	if err != nil {
		http.Error(writer, "failed to create user", http.StatusInternalServerError)
		return
	}

	newUserId, err := handler.dbContext.InsertUser(context.Background(), newUser)
	if err != nil {
		http.Error(writer, "failed to create user", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	http.ResponseWriter.WriteHeader(writer, http.StatusCreated)
	http.ResponseWriter.Write(writer, []byte("{\"message\": \"user created successfully\", \"userId\": "))
	http.ResponseWriter.Write(writer, []byte(fmt.Sprintf("%d", newUserId)))
	// for debugging
	http.ResponseWriter.Write(writer, []byte(", \"user\": "))
	userJson, _ := json.Marshal(newUser)
	http.ResponseWriter.Write(writer, userJson)
	http.ResponseWriter.Write(writer, []byte("}"))
}
