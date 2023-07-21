// Package zapx provides utilities for zapx.
package zapx

import (
	"fmt"

	"go.uber.org/zap"
)

// Config is zap logger configuration.
type Config struct {
	UseProductionLogger bool `envconfig:"USE_PRODUCTION_LOGGER"`
}

// NewLogger create instance of zap.
func NewLogger(config *Config) (*zap.Logger, error) {
	if config.UseProductionLogger {
		logger, err := zap.NewProduction()
		if err != nil {
			return nil, fmt.Errorf("create zapx error: %w", err)
		}
		return logger, nil
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, fmt.Errorf("create zapx error: %w", err)
	}
	return logger, nil
}
