package temporal

import (
	"go.uber.org/zap"
	"temporal-go-lib/internal/logging"
)

// NewLogger creates a new zap.Logger instance
func NewLogger() (*zap.Logger, error) {
	return logging.InitLogger()
}

// NewDevelopmentLogger creates a new zap.Logger instance for development
func NewDevelopmentLogger() (*zap.Logger, error) {
	return logging.InitDevelopmentLogger()
}
