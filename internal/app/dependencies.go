package app

import (
	"github.com/mace-codes/moondown-api/internal/app/config"
	"go.uber.org/zap"

	transport "github.com/mace-codes/go-keelson-transport"
)

// startable defines an interface for components that can be started. It requires a Start method that returns an error if the component fails to start.
type startable interface {
	Start() error
}

// stoppable defines an interface for components that can be stopped. It requires a Stop method that returns an error if the component fails to stop.
type stoppable interface {
	Stop() error
}

// dependencies holds the application's dependencies, such as database connections, external services, etc.
type dependencies struct {
	cfg    *config.Config
	logger *zap.Logger

	startables []startable
	stopables  []stoppable
}

// Critical returns a list of service critical health reporters
func (d *dependencies) Critical() []transport.Reporter {
	return []transport.Reporter{}
}

// Optional returns a list of service optional health reporters
func (d *dependencies) Optional() []transport.Reporter {
	return []transport.Reporter{}
}

// initializeDependencies sets up the necessary dependencies for the application based on the provided configuration.
func initializeDependencies(cfg *config.Config) (*dependencies, error) {

	// Build and initialize your dependencies here based on the configuration.

	return &dependencies{
		cfg:    cfg,
		logger: zap.NewExample(),

		// Initialize your dependencies here based on the configuration.

		startables: []startable{},
		stopables:  []stoppable{},
	}, nil
}
