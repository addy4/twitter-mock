package utils

import (
	"net/http"
	"strings"
)

func Authorize(requestHandler http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		r.Header.Set("Tw-Client-ID", "null")

		resp := ValidateToken(token)

		if resp.Active == false {
			w.WriteHeader(401)
		}

		r.Header.Set("Tw-Client-ID", resp.ClientId)
		requestHandler.ServeHTTP(w, r)
	})

}
