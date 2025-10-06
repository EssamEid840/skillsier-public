package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

type Claims struct {
	Email             string   `json:"email"`
	Sub               string   `json:"sub"`
	PreferredUsername string   `json:"preferred_username"`
	RealmAccess       struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
}

// Middleware verifies a Bearer JWT issued by `issuer` and intended for `clientID` (audience).
func Middleware(issuer, clientID string) (func(http.Handler) http.Handler, error) {
	provider, err := oidc.NewProvider(context.Background(), issuer)
	if err != nil {
		return nil, err
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authz := r.Header.Get("Authorization")
			if !strings.HasPrefix(authz, "Bearer ") {
				http.Error(w, "missing bearer", http.StatusUnauthorized)
				return
			}
			raw := strings.TrimPrefix(authz, "Bearer ")

			idt, err := verifier.Verify(r.Context(), raw)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			var c Claims
			if err := idt.Claims(&c); err != nil {
				http.Error(w, "bad claims", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), "claims", c)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}, nil
}
