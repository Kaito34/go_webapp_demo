package main

import (
	"net/http"
	"github.com/example/go_webapp_demo/pkg/render"
)

// Home is the home page hander
func Home(w http.ResponseWriter, r *http.Request){
	RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request){
	RenderTemplate(w, "about.page.tmpl")
}


