package temporal

import (
    "temporal-go-lib/internal/worker"
    "go.temporal.io/sdk/client"
    "time"
    "go.uber.org/zap"
)

// Worker wraps the internal worker
type Worker struct {
    internalWorker *worker.Worker
    logger         *zap.Logger
}

// WorkerOptions holds configuration for the Temporal worker
type WorkerOptions struct {
    TaskQueue                         string
    MaxConcurrentActivityExecutionSize int
    MaxConcurrentWorkflowTaskExecutionSize int
    WorkerActivitiesPerSecond          float64
    WorkerStopTimeout                  time.Duration
    Logger                             *zap.Logger
}

// NewWorker creates a new Temporal worker
func NewWorker(client *Client, options WorkerOptions) *Worker {
    internalWorker := worker.NewWorker(client.internalClient, worker.WorkerOptions{
        TaskQueue:                         options.TaskQueue,
        MaxConcurrentActivityExecutionSize: options.MaxConcurrentActivityExecutionSize,
        MaxConcurrentWorkflowTaskExecutionSize: options.MaxConcurrentWorkflowTaskExecutionSize,
        WorkerActivitiesPerSecond:          options.WorkerActivitiesPerSecond,
        WorkerStopTimeout:                  options.WorkerStopTimeout,
        Logger:                             options.Logger,
    })
    options.Logger.Info("Temporal worker created successfully")
    return &Worker{internalWorker: internalWorker, logger: options.Logger}
}

// RegisterWorkflow registers a workflow
func (w *Worker) RegisterWorkflow(workflow interface{}) {
    w.internalWorker.RegisterWorkflow(workflow)
    w.logger.Info("Workflow registered successfully")
}

// RegisterActivity registers an activity
func (w *Worker) RegisterActivity(activity interface{}) {
    w.internalWorker.RegisterActivity(activity)
    w.logger.Info("Activity registered successfully")
}

// Start starts the worker
func (w *Worker) Start() error {
    w.logger.Info("Starting Temporal worker...")
    err := w.internalWorker.Start()
    if err != nil {
        w.logger.Fatal("Failed to start worker", zap.Error(err))
        return err
    }
    w.logger.Info("Temporal worker started successfully")
    return nil
}

// Stop stops the worker
func (w *Worker) Stop() {
    w.internalWorker.Stop()
    w.logger.Info("Temporal worker stopped successfully")
}
