package workflows

import (
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	"time"
)

// BaseWorkflow provides common functionality for all workflows
type BaseWorkflow struct {
	WorkflowID     string
	RunID          string
	WorkflowLogger log.Logger
}

// Initialize initializes the workflow
func (b *BaseWorkflow) Initialize(ctx workflow.Context) {
	info := workflow.GetInfo(ctx)
	b.WorkflowID = info.WorkflowExecution.ID
	b.RunID = info.WorkflowExecution.RunID
	b.WorkflowLogger = workflow.GetLogger(ctx)
	b.WorkflowLogger.Info("Workflow initialized", zap.String("WorkflowID", b.WorkflowID), zap.String("RunID", b.RunID))
}

// SetSignalHandler sets a signal handler for the workflow
func (b *BaseWorkflow) SetSignalHandler(ctx workflow.Context, signalName string, handler interface{}) {
	b.SetSignalHandler(ctx, signalName, handler)
	b.WorkflowLogger.Info("Signal handler set", zap.String("SignalName", signalName))
}

// Sleep sleeps for the specified duration
func (b *BaseWorkflow) Sleep(ctx workflow.Context, duration time.Duration) error {
	b.WorkflowLogger.Info("Workflow sleeping", zap.Duration("Duration", duration))
	err := workflow.Sleep(ctx, duration)
	if err != nil {
		b.WorkflowLogger.Error("Error during workflow sleep", zap.Error(err))
		return err
	}
	b.WorkflowLogger.Info("Workflow woke up from sleep")
	return nil
}

// GetWorkflowInfo returns workflow information
func (b *BaseWorkflow) GetWorkflowInfo(ctx workflow.Context) *workflow.Info {
	info := workflow.GetInfo(ctx)
	b.WorkflowLogger.Info("Retrieved workflow info", zap.Any("Info", info))
	return info
}

// Logger WorkflowLogger returns a logger for the workflow
func (b *BaseWorkflow) Logger(ctx workflow.Context) log.Logger {
	return workflow.GetLogger(ctx)
}
