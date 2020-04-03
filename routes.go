package main

import (
	"github.com/go-chi/chi"
)

func getRoutes() chi.Router {
	// We're using chi as the router. You'll want to read
	// the documentation https://github.com/go-chi/chi

	r := chi.NewRouter()
	r.Get("/", controllerlanding)
	r.Get("/welcome", controllerwelcome)
	r.Get("/what_to_expect", controllerexpect)
	r.Get("/meet_the_team", controllermeet)
	r.Get("/contact", controllercontact)
	addStaticFileServer(r, "/static/", "staticfiles")
	return r
}
