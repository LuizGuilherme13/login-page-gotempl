package controllers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/repository"
)

func Signin(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username/e-mail")
	password := r.FormValue("password")

	user, err := repository.AuthWithUsernameOrEmail(username, password)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("Usuário e/ou senha incorretos!")
		http.Redirect(w, r, "/sigin", http.StatusSeeOther)
		return
	case err != nil:
		fmt.Println(err)
		http.Redirect(w, r, "/sigin", http.StatusSeeOther)
		return

	}

	ctx := context.WithValue(r.Context(), "user", user)
	r = r.WithContext(ctx)

	Home(w, r)
}
