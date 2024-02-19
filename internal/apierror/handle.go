package apierror

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, err error, userMessage string, statusCode int) {
	fmt.Println(err)
	http.Error(w, userMessage, statusCode)
}
