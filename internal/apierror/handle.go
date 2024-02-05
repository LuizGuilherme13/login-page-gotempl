package apierror

import (
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, err error, userMessage, internalMessage string, statusCode int) {

	log.Printf("controllers.Signup(): %s : %v", internalMessage, err)

	w.WriteHeader(statusCode)

	w.Write([]byte(userMessage))

}
