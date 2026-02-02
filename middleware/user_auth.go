package custommiddleware

import (
	"context"
	//"fmt"
	"net/http"

	"github.com/AbanoubGirges/malaykaproject/services"
)

func UserAuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("in middleware")
		token := r.Header.Get("Authorization")
		//fmt.Println("token:", token)
		c, err := services.ValidateJWT(token, services.SecretKey)
		if err != nil {
			services.RespondWithJson(w, 401, map[string]string{"error": "INVALID_TOKEN"})
			return
		}
		//claims := c.Claims.(jwt.MapClaims)

		claims := map[string]interface{}{}
		for k, v := range c {
		   claims[k] = v
		}
		if claims["user_id"] == nil {
			services.RespondWithJson(w, 401, map[string]string{"error": "MISSING_ID_IN_TOKEN"})
			return
		}

		//fmt.Println("validated token")
		// Store claims in context
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
		//fmt.Println("after next")
	})
}
