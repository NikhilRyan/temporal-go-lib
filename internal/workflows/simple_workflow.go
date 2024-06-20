package workflows

import (
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	"temporal-go-lib/internal/activities"
	"temporal-go-lib/internal/monitoring"
	"time"
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
		baseWorkflow.WorkflowLogger.Error("SimpleActivity failed", zap.Error(err))
		monitoring.WorkflowErrors.Inc()
		return err
	}

	baseWorkflow.WorkflowLogger.Info("SimpleActivity completed successfully", zap.String("Result", result))
	monitoring.WorkflowSuccess.Inc()
	return nil
}
