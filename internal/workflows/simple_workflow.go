package workflows

import (
    "go.temporal.io/sdk/workflow"
    "time"
    "temporal-go-lib/internal/activities"
    "go.uber.org/zap"
    "temporal-go-lib/internal/monitoring"
)

// SimpleWorkflow defines a simple workflow
func SimpleWorkflow(ctx workflow.Context) error {
    // Set activity options
    ao := workflow.ActivityOptions{
        StartToCloseTimeout: time.Minute,
        RetryPolicy: &temporal.RetryPolicy{
            InitialInterval:    time.Second,
            BackoffCoefficient: 2.0,
            MaximumInterval:    time.Minute,
            MaximumAttempts:    5,
        },
    }
    ctx = workflow.WithActivityOptions(ctx, ao)

    // Initialize the base workflow
    baseWorkflow := BaseWorkflow{}
    baseWorkflow.Initialize(ctx)

    var result string
    err := workflow.ExecuteActivity(ctx, activities.SimpleActivity, "Temporal").Get(ctx, &result)
    if err != nil {
        baseWorkflow.Logger.Error("SimpleActivity failed", zap.Error(err))
        monitoring.WorkflowErrors.Inc()
        return err
    }

    baseWorkflow.Logger.Info("SimpleActivity completed successfully", zap.String("Result", result))
    monitoring.WorkflowSuccess.Inc()
    return nil
}
