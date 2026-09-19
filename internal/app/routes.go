package app

import (
	"net/http"

	transport "github.com/mace-codes/go-keelson-transport"
)

// serveReach provides a simple response checkthat returns a JSON response indicating the service is accessable.
func serveReach() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transport.RespondJSON(w, r, http.StatusOK, nil, map[string]string{"status": "ok"})
	}
}

// serveRoutes takes in route groups exported by the rest interface of bounded contexts
func serveRoutes(deps *dependencies) ([]transport.RoutesConfig, error) {
	routes := []transport.RoutesConfig{
		{
			Path:     "/check-reach",
			Methods:  []string{http.MethodGet},
			Handler:  serveReach(),
			Adapters: nil,
		},
	}

	return routes, nil
}
