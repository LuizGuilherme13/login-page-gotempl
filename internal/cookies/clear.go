package cookies

import (
	"net/http"
	"time"
)

func Clear(w http.ResponseWriter) error {
	cookie := &http.Cookie{
		Name:     "user",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Expires:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}

	err := write(w, cookie)
	if err != nil {
		return err
	}
	return nil
}
