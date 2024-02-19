package controllers

import (
	"log"
	"net/http"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/cookies"
)

func LogOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	err := cookies.Clear(w)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
