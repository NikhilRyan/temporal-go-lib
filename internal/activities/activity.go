package activities

import (
    "context"
    "fmt"
    "go.uber.org/zap"
    "temporal-go-lib/internal/monitoring"
)

// SimpleActivity is an example of a simple activity
func SimpleActivity(ctx context.Context, name string) (string, error) {
    result := fmt.Sprintf("Hello, %s!", name)
    zap.L().Info("SimpleActivity executed", zap.String("Result", result))
    monitoring.ActivitySuccess.Inc()
    return result, nil
}

// ComplexActivity1 is an example of a complex activity
func ComplexActivity1(ctx context.Context, step string) (string, error) {
    // Simulate some processing
    zap.L().Info("ComplexActivity1 started", zap.String("Step", step))
    result := fmt.Sprintf("ComplexActivity1 completed for %s", step)
    zap.L().Info("ComplexActivity1 executed", zap.String("Result", result))
    monitoring.ActivitySuccess.Inc()
    return result, nil
}

// ComplexActivity2 is another example of a complex activity
func ComplexActivity2(ctx context.Context, step string) (string, error) {
    // Simulate some processing
    zap.L().Info("ComplexActivity2 started", zap.String("Step", step))
    result := fmt.Sprintf("ComplexActivity2 completed for %s", step)
    zap.L().Info("ComplexActivity2 executed", zap.String("Result", result))
    monitoring.ActivitySuccess.Inc()
    return result, nil
}
