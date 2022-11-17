package lib

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/spf13/viper"
	"github.com/uptrace/bunrouter"
)

type AuthMiddleware struct {
	oidc     *oidc.Provider
	verifier *oidc.IDTokenVerifier
}

func NewAuthMiddleware(ctx context.Context) (*AuthMiddleware, error) {
	p, err := oidc.NewProvider(ctx, viper.GetString("issuer_url"))
	if err != nil {
		return nil, err
	}

	v := p.Verifier(&oidc.Config{
		ClientID: viper.GetString("client_id"),
	})

	return &AuthMiddleware{
		oidc:     p,
		verifier: v,
	}, nil
}

func (a *AuthMiddleware) Middleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		h := req.Header.Get("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			return nil
		}

		t := strings.Replace(h, "Bearer ", "", 1)
		tok, err := a.verifier.Verify(req.Context(), t)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return nil
		}

		var claims struct {
			Name string `json:"name"`
		}
		if err := tok.Claims(&claims); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return nil
		}

		ctx := context.WithValue(
			req.Context(),
			"user_name",
			claims.Name,
		)

		return next(w, req.WithContext(ctx))
	}
}
