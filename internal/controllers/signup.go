package controllers

import (
	"context"
	"errors"
	"net/http"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/apierror"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/initializers"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/models"
	"github.com/LuizGuilherme13/norm/pkg/norm"
	"github.com/LuizGuilherme13/norm/pkg/norm/repository/postgres"
	"github.com/lib/pq"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	user.UserName = r.FormValue("username")
	user.Email = r.FormValue("e-mail")
	user.Password = r.FormValue("password")

	nORM := norm.NewService(postgres.New(initializers.DB))

	err := nORM.InTable("users").FromModel(user).Omit("user_id").Create()

	var pErr *pq.Error
	switch {
	case errors.As(err, &pErr) && pErr.Code == pq.ErrorCode("23505"):
		apierror.Handle(w, err, "This username is already in use",
			err.Error(), http.StatusBadRequest)
		return
	case err != nil:
		apierror.Handle(w, err, "An error occurred in the application",
			err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.WithValue(r.Context(), "user", user)
	r = r.WithContext(ctx)

	Home(w, r)
}
