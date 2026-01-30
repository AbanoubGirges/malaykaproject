package custommiddleware

import (
	"context"
	"net/http"

	"github.com/AbanoubGirges/malaykaproject/services"
)

func UserAuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authentication")
		claims, err := services.ValidateJWT(token, services.SecretKey)
		if err != nil {
			services.RespondWithJson(w, 401, struct{ error string }{error: "INVALID_TOKEN"})
			return
		}

		// Store claims in context
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
