package controllers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/apierror"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/cookies"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/repository"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username/e-mail")
	password := r.FormValue("password")

	authenticatedUser, err := repository.AuthWithUsernameOrEmail(username, password)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		apierror.Handle(w, err, "Incorrect username and/or password!", http.StatusUnauthorized)
		return

	case err != nil:
		apierror.Handle(w, err, "An error occurred in the application!", http.StatusInternalServerError)
		return
	}

	cookies.Set(w, authenticatedUser)

	http.Redirect(w, r, "/home", http.StatusFound)
}
