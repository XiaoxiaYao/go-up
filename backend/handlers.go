package main

import (
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	user, err := app.DB.GetUserByEmail("admin@example.com")
	if err != nil {
		fmt.Fprint(w, "this is the home page")
	}
	fmt.Fprint(w, user.Email)
}
