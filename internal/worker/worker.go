package worker

import (
    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
    "go.uber.org/zap"
    "temporal-go-lib/internal/monitoring"
    "time"
)

// WorkerOptions holds configuration for the Temporal worker
type WorkerOptions struct {
    TaskQueue           string
    MaxConcurrentActivityExecutionSize int
    MaxConcurrentWorkflowTaskExecutionSize int
    WorkerActivitiesPerSecond float64
    WorkerStopTimeout         time.Duration
    Logger                    *zap.Logger
}

// Worker wraps the Temporal SDK worker
type Worker struct {
    temporalWorker worker.Worker
    logger         *zap.Logger
}

// NewWorker creates a new Temporal worker
func NewWorker(client client.Client, options WorkerOptions) *Worker {
    w := worker.New(client, options.TaskQueue, worker.Options{
        MaxConcurrentActivityExecutionSize: options.MaxConcurrentActivityExecutionSize,
        MaxConcurrentWorkflowTaskExecutionSize: options.MaxConcurrentWorkflowTaskExecutionSize,
        WorkerActivitiesPerSecond: options.WorkerActivitiesPerSecond,
        WorkerStopTimeout: options.WorkerStopTimeout,
    })
    options.Logger.Info("Temporal worker created successfully")
    return &Worker{temporalWorker: w, logger: options.Logger}
}

// RegisterWorkflow registers a workflow with the worker
func (w *Worker) RegisterWorkflow(workflow interface{}) {
    w.temporalWorker.RegisterWorkflow(workflow)
    w.logger.Info("Workflow registered successfully")
}

// RegisterActivity registers an activity with the worker
func (w *Worker) RegisterActivity(activity interface{}) {
    w.temporalWorker.RegisterActivity(activity)
    w.logger.Info("Activity registered successfully")
}

// Start starts the Temporal worker
func (w *Worker) Start() error {
    w.logger.Info("Starting Temporal worker...")
    err := w.temporalWorker.Start()
    if err != nil {
        w.logger.Fatal("Failed to start worker", zap.Error(err))
        return err
    }
    monitoring.WorkerStart.Inc()
    w.logger.Info("Temporal worker started successfully")
    return nil
}

// Stop stops the Temporal worker
func (w *Worker) Stop() {
    w.temporalWorker.Stop()
    monitoring.WorkerStop.Inc()
    w.logger.Info("Temporal worker stopped successfully")
}
