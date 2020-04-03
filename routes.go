package main

import (
	"github.com/go-chi/chi"
)

func getRoutes() chi.Router {
	// We're using chi as the router. You'll want to read
	// the documentation https://github.com/go-chi/chi

	r := chi.NewRouter()
	r.Get("/", controllerlanding)
	r.get("/welcome", controllerwelcome)
	r.get("/what_to_expect", controllerexpect)
	r.get("/meet_the_team", controllermeet)
	r.get("/contact", controllercontact)
	addStaticFileServer(r, "/static/", "staticfiles")
	return r
}
