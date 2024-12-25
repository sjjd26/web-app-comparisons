package AuthenticationHandlers

import (
	"fmt"
	"net/http"
)

func RegisterNewUser(writer http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	password := request.FormValue("password")

	if email == "" || password == "" {
		http.Error(writer, "Missing email or password", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(writer, "Username: %s, Password: %s", email, password)
}
