package custommiddleware

import (
	"fmt"
	"net/http"
	//"github.com/AbanoubGirges/malaykaproject/controllers"
	"github.com/AbanoubGirges/malaykaproject/services"
)

func AdminAuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement your admin authentication logic here
		// For example, check for a specific header or token
		token := r.Header.Get("Authorization")
		claims, err := services.ValidateJWT(token, services.SecretKey)
		if err != nil {
			services.RespondWithJson(w, 403, map[string]string{"error": "INVALID_TOKEN"})
			fmt.Print("invalid token")
			return
		}
		claimsRole, ok := claims["role"]
		if !ok || claimsRole != "admin" {
			services.RespondWithJson(w, 403, map[string]string{"error": "INVALID_CREDS"})
			fmt.Print("invalid creds")
			return
		}
		next.ServeHTTP(w, r)
	})
}
