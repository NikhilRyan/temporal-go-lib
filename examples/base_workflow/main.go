package main

import (
	"go.uber.org/zap"
	"log"
	"temporal-go-lib/internal/activities"
	"temporal-go-lib/internal/monitoring"
	"temporal-go-lib/internal/workflows"
	"temporal-go-lib/pkg/temporal"
	"time"
)

func main() {
	// Initialize logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	// Start monitoring
	go monitoring.ServeMetrics()

	// Create Temporal client
	client, err := temporal.NewClient(temporal.ClientOptions{
		HostPort:  "localhost:7233",
		Namespace: "default",
		Logger:    logger,
	})
	if err != nil {
		logger.Fatal("Failed to create Temporal client", zap.Error(err))
	}
	defer client.Close()

	// Worker options
	workerOptions := temporal.WorkerOptions{
		TaskQueue:                              "base-task-queue",
		MaxConcurrentActivityExecutionSize:     10,
		MaxConcurrentWorkflowTaskExecutionSize: 10,
		WorkerActivitiesPerSecond:              100.0,
		WorkerStopTimeout:                      time.Minute,
		Logger:                                 logger,
	}

	// Create and start worker
	worker := temporal.NewWorker(client, workerOptions)
	worker.RegisterWorkflow(workflows.BaseWorkflow{}.Initialize)
	worker.RegisterActivity(activities.SimpleActivity)

	err = worker.Start()
	if err != nil {
		logger.Fatal("Failed to start Temporal worker", zap.Error(err))
	}
	defer worker.Stop()

	logger.Info("Worker started successfully")
}
