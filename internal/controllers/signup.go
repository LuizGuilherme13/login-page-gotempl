package controllers

import (
	"database/sql"
	"net/http"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/apierror"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/initializers"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/models"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/repository"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("e-mail")
	password := r.FormValue("password")

	db, err := sql.Open("postgres", initializers.DB.String())
	if err != nil {
		apierror.Handle(w, err, "Ocorreu um erro na aplicação",
			"Error opening database", http.StatusInternalServerError)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		apierror.Handle(w, err, "Ocorreu um erro na aplicação",
			"Error starting transaction", http.StatusInternalServerError)
	}
	defer tx.Rollback()

	// query := "INSERT INTO users(username, email, password) "
	// query += "VALUES($1, $2, $3)"

	// _, err = tx.Exec(query, username, email, password)
	// if err != nil {
	// 	apierror.Handle(w, err, "Ocorreu um erro na aplicação",
	// 		"Error when executing", http.StatusInternalServerError)
	// }

	err = repository.CreateUser(tx, models.User{
		UserName: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		apierror.Handle(w, err, "Ocorreu um erro na aplicação", err.Error(), http.StatusInternalServerError)
	}

	if err := tx.Commit(); err != nil {
		apierror.Handle(w, err, "Ocorreu um erro na aplicação",
			"Error committing transaction", http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}
