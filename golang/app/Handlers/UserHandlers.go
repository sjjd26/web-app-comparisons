package handlers

import (
	"fmt"
	"net/http"
	databasecontext "web-app/auth-api/DatabaseContext"
)

type UserHandler struct {
	databaseContext databasecontext.DatabaseContext
}

func NewUserHandler(databaseContext databasecontext.DatabaseContext) *UserHandler {
	return &UserHandler{
		databaseContext: databaseContext,
	}
}

func (handler UserHandler) RegisterNewUser(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")

	if email == "" || password == "" {
		http.Error(writer, "Missing email or password", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(writer, "Username: %s, Password: %s", email, password)
}
