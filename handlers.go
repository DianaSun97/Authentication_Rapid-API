package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + templateName)
	if err != nil {
		fmt.Println("error Parse", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error Execute", err)
		return
	}
}

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home-page.gohtml")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about-page.gohtml")
}
