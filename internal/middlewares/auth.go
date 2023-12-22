package middlewares

import "net/http"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")

		if user == nil {
			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
