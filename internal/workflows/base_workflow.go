package workflows

import (
    "go.temporal.io/sdk/workflow"
    "go.uber.org/zap"
    "time"
)

// BaseWorkflow provides common functionality for all workflows
type BaseWorkflow struct {
    WorkflowID string
    RunID      string
    Logger     *zap.Logger
}

// Initialize initializes the workflow
func (b *BaseWorkflow) Initialize(ctx workflow.Context) {
    info := workflow.GetInfo(ctx)
    b.WorkflowID = info.WorkflowExecution.ID
    b.RunID = info.WorkflowExecution.RunID
    b.Logger = workflow.GetLogger(ctx)
    b.Logger.Info("Workflow initialized", zap.String("WorkflowID", b.WorkflowID), zap.String("RunID", b.RunID))
}

// SetSignalHandler sets a signal handler for the workflow
func (b *BaseWorkflow) SetSignalHandler(ctx workflow.Context, signalName string, handler interface{}) {
    workflow.SetSignalHandler(ctx, signalName, handler)
    b.Logger.Info("Signal handler set", zap.String("SignalName", signalName))
}

// Sleep sleeps for the specified duration
func (b *BaseWorkflow) Sleep(ctx workflow.Context, duration time.Duration) error {
    b.Logger.Info("Workflow sleeping", zap.Duration("Duration", duration))
    err := workflow.Sleep(ctx, duration)
    if err != nil {
        b.Logger.Error("Error during workflow sleep", zap.Error(err))
        return err
    }
    b.Logger.Info("Workflow woke up from sleep")
    return nil
}

// GetWorkflowInfo returns workflow information
func (b *BaseWorkflow) GetWorkflowInfo(ctx workflow.Context) workflow.Info {
    info := workflow.GetInfo(ctx)
    b.Logger.Info("Retrieved workflow info", zap.Any("Info", info))
    return info
}

// Logger returns a logger for the workflow
func (b *BaseWorkflow) Logger(ctx workflow.Context) workflow.Logger {
    return workflow.GetLogger(ctx)
}
