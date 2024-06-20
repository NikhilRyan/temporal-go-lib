package temporal

import (
    "temporal-go-lib/internal/activities"
    "context"
    "go.uber.org/zap"
)

// SimpleActivity is an example of a simple activity
func SimpleActivity(ctx context.Context, name string) (string, error) {
    result, err := activities.SimpleActivity(ctx, name)
    if err != nil {
        zap.L().Error("SimpleActivity failed", zap.Error(err))
        return "", err
    }
    zap.L().Info("SimpleActivity executed successfully", zap.String("Result", result))
    return result, nil
}

// ComplexActivity1 is an example of a complex activity
func ComplexActivity1(ctx context.Context, step string) (string, error) {
    result, err := activities.ComplexActivity1(ctx, step)
    if err != nil {
        zap.L().Error("ComplexActivity1 failed", zap.Error(err))
        return "", err
    }
    zap.L().Info("ComplexActivity1 executed successfully", zap.String("Result", result))
    return result, nil
}

// ComplexActivity2 is another example of a complex activity
func ComplexActivity2(ctx context.Context, step string) (string, error) {
    result, err := activities.ComplexActivity2(ctx, step)
    if err != nil {
        zap.L().Error("ComplexActivity2 failed", zap.Error(err))
        return "", err
    }
    zap.L().Info("ComplexActivity2 executed successfully", zap.String("Result", result))
    return result, nil
}
