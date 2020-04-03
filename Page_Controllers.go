package main

import (
	"net/http"
)

func controllerlanding(w http.ResponseWriter, r *http.Request) {

	tmpl["landing"].Execute(w, nil)
}

func controllerwelcome(w http.ResponseWriter, r *http.Request) {

	tmpl["welcome"].Execute(w, nil)
}

func controllerexpect(w http.ResponseWriter, r *http.Request) {

	tmpl["what_to_expect"].Execute(w, nil)
}

func controllermeet(w http.ResponseWriter, r *http.Request) {

	tmpl["meet_the_team"].Execute(w, nil)
}

func controllercontact(w http.ResponseWriter, r *http.Request) {

	tmpl["contact_us"].Execute(w, nil)
}
