package handlers

import (
	"net/http"

	"github.com/Kaito34/go_webapp_demo/pkg/render"
)

// Home is the home page hander
func Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "about.page.tmpl")
}


