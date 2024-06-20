package temporal

import (
    "go.uber.org/zap"
    "temporal-go-lib/internal/errorhandling"
)

// HandleWorkflowError handles errors in workflow execution
func HandleWorkflowError(logger *zap.Logger, err error) error {
    return errorhandling.HandleWorkflowError(logger, err)
}

// HandleActivityError handles errors in activity execution
func HandleActivityError(logger *zap.Logger, err error) error {
    return errorhandling.HandleActivityError(logger, err)
}

// HandleClientError handles errors in client operations
func HandleClientError(logger *zap.Logger, err error) error {
    return errorhandling.HandleClientError(logger, err)
}
