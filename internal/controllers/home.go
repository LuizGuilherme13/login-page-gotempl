package controllers

import (
	"net/http"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/models"
	"github.com/LuizGuilherme13/login-page-gotempl/internal/views"
)

func Home(w http.ResponseWriter, r *http.Request) {

	authenticatedUser, ok := r.Context().Value("user").(*models.User)
	if !ok {
		http.Error(w, "Usuário não encontrado no contexto", http.StatusInternalServerError)
		return
	}

	views.Home(*authenticatedUser).Render(r.Context(), w)

}
