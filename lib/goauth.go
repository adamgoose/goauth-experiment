package lib

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

type GoauthServer struct {
	http.Handler
	Auth *AuthMiddleware
}

func NewGoauthServer(ctx context.Context) (*GoauthServer, error) {
	r := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
	)
	a, err := NewAuthMiddleware(ctx)
	if err != nil {
		return nil, err
	}

	g := &GoauthServer{
		Handler: r,
		Auth:    a,
	}

	r.Compat().GET("/", g.Index)

	r.Use(a.Middleware).Compat().WithGroup("/api", func(r *bunrouter.CompatGroup) {
		r.GET("/secret", g.Secret)
	})

	return g, nil
}

func (g *GoauthServer) Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"challenge": "Do it right!",
	})
}

func (g *GoauthServer) Secret(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"prize":     "satisfaction",
		"recipient": r.Context().Value("user_name").(string),
	})
}
