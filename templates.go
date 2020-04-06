package main

import (
	"html/template"
)

var tmpl = make(map[string]*template.Template)

func init() {
	m := template.Must
	p := template.ParseFiles
	tmpl["landing"] = m(p("templates/landing.gohtml", "templates/layout.gohtml"))
	tmpl["welcome"] = m(p("templates/welcome.gohtml", "templates/layout.gohtml"))
	tmpl["investment_criteria"] = m(p("templates/investment_criteria.gohtml", "templates/layout.gohtml"))
	tmpl["what_to_expect"] = m(p("templates/what_to_expect.gohtml", "templates/layout.gohtml"))
	tmpl["meet_the_team"] = m(p("templates/meet_the_team.gohtml", "templates/layout.gohtml"))
	tmpl["contact_us"] = m(p("templates/contact_us.gohtml", "templates/layout.gohtml"))
}
