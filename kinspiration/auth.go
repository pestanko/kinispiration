package kinspiration

import (
	"fmt"
	"log"
	"net/http"
)

type AuthMiddleware struct {
	App *App
}

func (auth *AuthMiddleware) Token() string {
	return auth.App.Config.AdminToken
}

func (auth *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		log.Printf("[AUTH] Received authorization header: %s", authorization)
		if r.Method == "GET" || authorization == "" {
			next.ServeHTTP(w, r)
			return
		}
		token := authorization

		if token == fmt.Sprintf("Bearer %s", auth.Token()) {
			// We found the token in our map
			log.Printf("[AUTH] Successfully authorized")
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
