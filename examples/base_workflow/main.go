package main

import (
    "log"
    "time"
    "go.uber.org/zap"
    "temporal-go-lib/pkg/temporal"
    "temporal-go-lib/pkg/temporal/worker"
    "temporal-go-lib/pkg/temporal/workflow"
    "temporal-go-lib/pkg/temporal/activity"
    "temporal-go-lib/internal/monitoring"
)

func main() {
    // Initialize logger
    logger, err := zap.NewProduction()
    if err != nil {
        log.Fatalf("Failed to create logger: %v", err)
    }
    defer logger.Sync()

    // Start monitoring
    go monitoring.ServeMetrics()

    // Create Temporal client
    client, err := temporal.NewClient(temporal.ClientOptions{
        HostPort: "localhost:7233",
        Namespace: "default",
        Logger: logger,
    })
    if err != nil {
        logger.Fatal("Failed to create Temporal client", zap.Error(err))
    }
    defer client.Close()

    // Worker options
    workerOptions := temporal.WorkerOptions{
        TaskQueue: "base-task-queue",
        MaxConcurrentActivityExecutionSize: 10,
        MaxConcurrentWorkflowTaskExecutionSize: 10,
        WorkerActivitiesPerSecond: 100.0,
        WorkerStopTimeout: time.Minute,
        Logger: logger,
    }

    // Create and start worker
    worker := temporal.NewWorker(client, workerOptions)
    worker.RegisterWorkflow(workflow.BaseWorkflow{}.Initialize)
    worker.RegisterActivity(activity.SimpleActivity)

    err = worker.Start()
    if err != nil {
        logger.Fatal("Failed to start Temporal worker", zap.Error(err))
    }
    defer worker.Stop()

    logger.Info("Worker started successfully")
}
