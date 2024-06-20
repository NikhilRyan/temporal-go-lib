package temporal

import (
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/workflow"
	"temporal-go-lib/internal/workflows"
	"time"
)

// BaseWorkflow provides common functionality for all workflows
type BaseWorkflow struct {
	workflows.BaseWorkflow
}

// SimpleWorkflow defines a simple workflow
func SimpleWorkflow(ctx workflow.Context) error {
	return workflows.SimpleWorkflow(ctx)
}

// ComplexWorkflow defines a complex workflow
func ComplexWorkflow(ctx workflow.Context) error {
	return workflows.ComplexWorkflow(ctx)
}

// Initialize initializes the workflow
func (b *BaseWorkflow) Initialize(ctx workflow.Context) {
	b.BaseWorkflow.Initialize(ctx)
}

// SetSignalHandler sets a signal handler for the workflow
func (b *BaseWorkflow) SetSignalHandler(ctx workflow.Context, signalName string, handler interface{}) {
	b.BaseWorkflow.SetSignalHandler(ctx, signalName, handler)
}

// Sleep sleeps for the specified duration
func (b *BaseWorkflow) Sleep(ctx workflow.Context, duration time.Duration) error {
	return b.BaseWorkflow.Sleep(ctx, duration)
}

// GetWorkflowInfo returns workflow information
func (b *BaseWorkflow) GetWorkflowInfo(ctx workflow.Context) *workflow.Info {
	return b.BaseWorkflow.GetWorkflowInfo(ctx)
}

// Logger returns a logger for the workflow
func (b *BaseWorkflow) Logger(ctx workflow.Context) log.Logger {
	return b.BaseWorkflow.Logger(ctx)
}
