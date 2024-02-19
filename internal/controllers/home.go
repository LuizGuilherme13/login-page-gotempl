package controllers

import (
	"context"
	"errors"
	"net/http"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/cookies"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/views"
)

func Home(w http.ResponseWriter, r *http.Request) {
	authenticatedUser, err := cookies.Get(w, r)
	switch {
	case errors.Is(err, http.ErrNoCookie):
		http.Error(w, "cookie not found", http.StatusBadRequest)
		return

	case errors.Is(err, cookies.ErrInvalidValue):
		http.Error(w, "invalid cookie", http.StatusBadRequest)
		return

	case err != nil:
		http.Error(w, "An error occurred in the application", http.StatusInternalServerError)
		return

	}

	views.Home(*authenticatedUser).Render(context.Background(), w)
}
