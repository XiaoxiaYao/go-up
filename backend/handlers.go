package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if app.Session.Exists(r.Context(), "test") {
		msg := app.Session.GetString(r.Context(), "test")
		fmt.Fprint(w, msg)
		return
	} else {
		app.Session.Put(r.Context(), "test", "hit this page"+time.Now().UTC().String())
	}
	fmt.Fprint(w, "this is the home page")
}
