package utils

import (
	"fmt"
	"net/http"
	"strings"
)

func Authorize(requestHandler http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		headers := r.Header
		fmt.Println("Headers of the WebSocket client request:")
		for key, value := range headers {
			fmt.Printf("%s: %s\n", key, value)
		}

		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if ValidateToken(token).Active == false {
			w.WriteHeader(401)
		}

		requestHandler.ServeHTTP(w, r)
	})

}
