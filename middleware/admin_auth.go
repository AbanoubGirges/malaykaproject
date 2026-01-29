package custommiddleware

import (
	"net/http"
	//"github.com/AbanoubGirges/malaykaproject/controllers"
	"github.com/AbanoubGirges/malaykaproject/services"
)

func AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement your admin authentication logic here
		// For example, check for a specific header or token
		token := r.Header.Get("Authorization")
		claims, err := services.ValidateJWT(token, services.SecretKey)
		if err != nil {
			services.RespondWithJson(w, 403, struct{ error string }{error: "INVALID_TOKEN"})
			return
		}
		claimsRole, ok := claims["role"].(uint)
		if !ok || claimsRole != 1 {
			services.RespondWithJson(w, 403, struct{ error string }{error: "INVALID_CREDS"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
