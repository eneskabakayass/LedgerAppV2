package middleware

import (
	"net/http"
)

func RequireRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := "admin"

			if userRole != role {
				http.Error(w, "Access Denied", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
