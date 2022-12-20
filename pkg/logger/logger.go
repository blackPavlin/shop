package logger

import "go.uber.org/zap"

// LoggerConfig
type LoggerConfig struct {
	UseProductionLogger bool `envconfig:"USE_PRODUCTION_LOGGER"`
}

// NewLogger
func NewLogger(config *LoggerConfig) (*zap.Logger, error) {
	if config.UseProductionLogger {
		return zap.NewDevelopment()
	}

	return zap.NewDevelopment()
}
