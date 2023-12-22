package main

import (
	"log"
	"net/http"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/controllers"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/initializers"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/middlewares"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/views"
	"github.com/a-h/templ"
)

func init() {
	initializers.LoadDB()
}

func main() {
	http.Handle("/", templ.Handler(views.Signin()))
	http.HandleFunc("/signin", middlewares.Logging(controllers.Signin))

	http.Handle("/signup_page", templ.Handler(views.Signup()))
	http.HandleFunc("/signup", middlewares.Logging(controllers.Signup))

	log.Println("Runnin on port :8080")
	http.ListenAndServe(":8080", nil)
}
