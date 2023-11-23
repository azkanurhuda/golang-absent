package middleware

import (
	"github.com/azkanurhuda/golang-absent/interfaces/presenter"
	"net/http"
	"strings"
)

func (h *Middleware) JWT() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			extractedToken := strings.Split(authHeader, "Bearer ")
			if len(extractedToken) != 2 {
				presenter.Error(w, http.StatusForbidden, "Invalid header")
				return
			}

			token := strings.TrimSpace(extractedToken[1])
			userID, err := h.jwtService.Verify(token)
			if err != nil {
				presenter.Error(w, http.StatusForbidden, "Invalid token")
				return
			}

			ctx := SetUserID(r.Context(), userID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		})
	}
}
