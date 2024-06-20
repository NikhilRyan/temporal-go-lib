package workflows

import (
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	"temporal-go-lib/internal/activities"
	"time"
)

// ComplexWorkflow defines a complex workflow
func ComplexWorkflow(ctx workflow.Context) error {
	// Set activity options
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 5,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 10,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute * 2,
			MaximumAttempts:    10,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Initialize the base workflow
	baseWorkflow := BaseWorkflow{}
	baseWorkflow.Initialize(ctx)

	// Execute first activity
	var result1 string
	err := workflow.ExecuteActivity(ctx, activities.ComplexActivity1, "Step 1").Get(ctx, &result1)
	if err != nil {
		baseWorkflow.WorkflowLogger.Error("ComplexActivity1 failed", zap.Error(err))
		return err
	}
	baseWorkflow.WorkflowLogger.Info("ComplexActivity1 completed successfully", zap.String("Result", result1))

	// Sleep for a duration
	err = baseWorkflow.Sleep(ctx, time.Minute)
	if err != nil {
		return err
	}

	// Execute second activity
	var result2 string
	err = workflow.ExecuteActivity(ctx, activities.ComplexActivity2, "Step 2").Get(ctx, &result2)
	if err != nil {
		baseWorkflow.WorkflowLogger.Error("ComplexActivity2 failed", zap.Error(err))
		return err
	}
	baseWorkflow.WorkflowLogger.Info("ComplexActivity2 completed successfully", zap.String("Result", result2))

	// Complete workflow
	baseWorkflow.WorkflowLogger.Info("ComplexWorkflow completed successfully")
	return nil
}
