package middlewares

import (
	"context"
	"famesensor/go-graphql-jwt/models"
	"famesensor/go-graphql-jwt/utils/jwt"
	"net/http"
)

type authString string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := jwt.JwtValidate(auth)
		if err != nil || !validate.Valid {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		customClaim, _ := validate.Claims.(*models.JwtCustomClaim)

		ctx := context.WithValue(r.Context(), authString("auth"), customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func GetJwtClaimFromHeader(ctx context.Context) *models.JwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*models.JwtCustomClaim)
	return raw
}
