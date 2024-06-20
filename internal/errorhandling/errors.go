package errorhandling

import (
    "errors"
    "go.uber.org/zap"
)

// Custom error definitions
var (
    ErrWorkflowFailed = errors.New("workflow execution failed")
    ErrActivityFailed = errors.New("activity execution failed")
    ErrClientFailure  = errors.New("client failure")
)

// HandleWorkflowError handles errors in workflow execution
func HandleWorkflowError(logger *zap.Logger, err error) error {
    logger.Error("Workflow error occurred", zap.Error(err))
    switch {
    case errors.Is(err, ErrWorkflowFailed):
        // Custom handling for workflow failure
        return err
    default:
        // General error handling
        return err
    }
}

// HandleActivityError handles errors in activity execution
func HandleActivityError(logger *zap.Logger, err error) error {
    logger.Error("Activity error occurred", zap.Error(err))
    switch {
    case errors.Is(err, ErrActivityFailed):
        // Custom handling for activity failure
        return err
    default:
        // General error handling
        return err
    }
}

// HandleClientError handles errors in client operations
func HandleClientError(logger *zap.Logger, err error) error {
    logger.Error("Client error occurred", zap.Error(err))
    switch {
    case errors.Is(err, ErrClientFailure):
        // Custom handling for client failure
        return err
    default:
        // General error handling
        return err
    }
}
