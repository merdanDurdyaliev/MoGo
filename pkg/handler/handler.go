package handler

import (
	"net/http"
)

func (app *Application) homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "not found", 404)
		return
	}
	app.authData(w, r)
	app.render(w, r, "index.html")
}

func (app *Application) storyAbout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "./static/storyAbout" {
		http.Error(w, "not found", 404)
		return
	}
	app.storyAbout(w, r)
	app.render(w, r, "storyAbout.html")
}

func (app *Application) ourServices(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "./static/services" {
		http.Error(w, "not found", 404)
		return
	}
	app.ourServices(w, r)
	app.render(w, r, "services.html")
}

func (app *Application) ourWorks(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "./static/works" {
		http.Error(w, "not found", 404)
		return
	}
	app.ourWorks(w, r)
	app.render(w, r, "works.html")
}

func (app *Application) ourBlogs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "./static/blogs" {
		http.Error(w, "not found", 404)
		return
	}
	app.ourBlogs(w, r)
	app.render(w, r, "works.html")
}

func (app *Application) ourContacts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "./static/contacts" {
		http.Error(w, "not found", 404)
		return
	}
	app.ourContacts(w, r)
	app.render(w, r, "contacts.html")
}
