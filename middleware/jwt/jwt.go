package jwt

import (
	"hanye/pkg/util"
	"net/http"
	"time"
)

type JwtAuthMiddleware struct {
}

func NewJwtAuthMiddleware() *JwtAuthMiddleware {
	return &JwtAuthMiddleware{}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			// If the Authorization header is missing, return an error
			http.Error(w, "Unauthorized: Missing Authorization header", http.StatusUnauthorized)
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			http.Error(w, "Unauthorized: Authorization header Format Error", http.StatusUnauthorized)
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			http.Error(w, "Token 过期重新登录", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
