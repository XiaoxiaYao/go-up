package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
)

func Test_app_routes(t *testing.T) {
	var registered = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
	}

	var app application
	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, route := range registered {

	}
}

func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false
}
