package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ctxKey string

const keyUserID ctxKey = "user_id"

func main() {
	mux := http.NewServeMux()

	mux.Handle("/restricted", authMiddleware(handleRestricted()))

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")

		if token != "very-secret-token" {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), keyUserID, 42)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleRestricted() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(keyUserID).(int)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "internal error, try again later please")
			return
		}

		io.WriteString(w, fmt.Sprintf("hello, user #%d!", userID))
	})
}
