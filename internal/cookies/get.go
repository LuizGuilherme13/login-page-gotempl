package cookies

import (
	"encoding/gob"
	"net/http"
	"strings"

	"github.com/LuizGuilherme13/login-page-gotempl/internal/models"
)

func Get(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	gobEncodedValue, err := read(r, "user")
	if err != nil {
		return nil, err
	}

	var user models.User

	/* Create an strings.Reader containing the gob-encoded value. */
	reader := strings.NewReader(gobEncodedValue)

	if err := gob.NewDecoder(reader).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
