package cookies

import (
	"bytes"
	"encoding/gob"
	"log"
	"net/http"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/models"
)

func Set(w http.ResponseWriter, authenticatedUser *models.User) {

	// err := Clear(w)
	// if err != nil {
	// 	log.Println(err)
	// 	http.Error(w, "server error", http.StatusInternalServerError)
	// 	return
	// }

	/* Initialize a buffer to hold the gob-encoded data. */
	var buf bytes.Buffer

	/* Gob-encode the user data, storing the encoded output in the buffer. */
	err := gob.NewEncoder(&buf).Encode(&authenticatedUser)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	/* Call buf.String() to get the gob-encoded value as a string and set it as
	the cookie value. */
	cookie := &http.Cookie{
		Name:     "user",
		Value:    buf.String(),
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	err = write(w, cookie)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

}
