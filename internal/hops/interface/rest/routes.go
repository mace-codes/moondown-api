package rest

import (
	transport "github.com/mace-codes/go-keelson-transport"
)

// Routes returns the HTTP routes for the hops bounded context.
func Routes() []transport.RoutesConfig {
	return []transport.RoutesConfig{
		{},
	}
}
