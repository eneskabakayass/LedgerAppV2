package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ValidateJSONBody(v interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Invalid body", http.StatusBadRequest)
				return
			}

			if err := json.Unmarshal(body, v); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
