package routes

import (
	"context"
	"log"
	"net/http"
	"slices"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func (app Dependencies[T]) EnsureValidToken(next http.Handler) http.Handler {
	app.Logger().Info("Ensure valid token hit")

	middleware := jwtmiddleware.New(
		app.JwtValidator().ValidateToken,
		jwtmiddleware.WithErrorHandler(func(w http.ResponseWriter, _ *http.Request, err error) {
			log.Printf("JWT validation error: %v", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message":"Invalid JWT"}`))
		}),
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware.CheckJWT(next).ServeHTTP(w, r)
	})
}

func (c CustomClaims) HasScope(expectedScope string) bool {
	result := strings.Split(c.Scope, " ")
	// for i := range result {
	// 	if result[i] == expectedScope {
	// 		return true
	// 	}
	// }
	return slices.Contains(result, expectedScope)
}
