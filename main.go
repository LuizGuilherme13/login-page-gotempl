package main

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/controllers"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/middlewares"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/models"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/views"
	"github.com/a-h/templ"
)

const SecretKey = "6d9e8faabd22cde539f5f3ae0fdc93d7aad06a038eca73a3ed1dea8a26b47153"

func main() {
	/* Importantly, we need to tell the encoding/gob package about the Go type
	that we want to encode. We do this my passing *an instance* of the type
	to gob.Register(). In this case we pass a pointer to an initialized (but
	empty) instance of the User struct. */
	gob.Register(&models.User{})

	http.Handle("/", templ.Handler(views.Signin()))
	http.HandleFunc("/signin", middlewares.Logging(controllers.Signin))
	http.HandleFunc("/home", middlewares.Logging(controllers.Home))
	http.HandleFunc("/logout", middlewares.Logging(controllers.LogOut))

	// http.Handle("/signup_page", templ.Handler(views.Signup()))
	// http.HandleFunc("/signup", middlewares.Logging(controllers.Signup))

	log.Println("Runnin on port :8080")
	http.ListenAndServe(":8080", nil)

}
