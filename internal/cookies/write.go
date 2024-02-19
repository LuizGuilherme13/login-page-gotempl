package cookies

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
)

var ErrValueToLong = errors.New("cookie value to long")

func write(w http.ResponseWriter, cookie *http.Cookie) error {
	// Encode the cookie value using base64.
	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))

	// Check the total length of the cookie contents. Return the ErrValueTooLong
	// error if it's more than 4096 bytes.
	if len(cookie.String()) > 4096 {
		return ErrValueToLong
	}

	// Write the cookie as normal.
	http.SetCookie(w, cookie)

	return nil
}

func writeSigned(w http.ResponseWriter, cookie http.Cookie, secretKey []byte) error {
	// Calculate a HMAC signature of the cookie name and value, using SHA256 and
	// a secret key (which we will create in a moment).
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(cookie.Name))
	mac.Write([]byte(cookie.Value))
	signature := mac.Sum(nil)

	// Prepend the cookie value with the HMAC signature.
	cookie.Value = string(signature) + cookie.Value

	// Call our Write() helper to base64-encode the new cookie value and write
	// the cookie.
	return write(w, &cookie)
}
