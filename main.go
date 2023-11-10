package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/michaelmagen/sync-spotify/routes"
)

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Routes
	r.Route("/auth", routes.AuthRoute)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}

// TODO: Handle env better, should only read them in one time. Make sure this will work in deploy as well.
// TODO: Make middleware for checking that user is authed
// TODO: Make sure deleting access tokens that are no longer needed. When auth middleware, if refresh the token, add new one then remove old one.
// TODO: If auth middleware fails then the front end should send the user to login
