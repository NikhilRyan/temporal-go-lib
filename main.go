package main

import (
	"go.uber.org/zap"
	"log"
	"temporal-go-lib/internal/config"
	"temporal-go-lib/internal/monitoring"
	"temporal-go-lib/pkg/temporal"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	var logger *zap.Logger
	if config.Logging.Level == "development" {
		logger, err = temporal.NewDevelopmentLogger()
	} else {
		logger, err = temporal.NewLogger()
	}
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	// Start monitoring if enabled
	if config.Monitoring.Enabled {
		go monitoring.ServeMetrics()
		logger.Info("Monitoring server started", zap.Int("port", config.Monitoring.Port))
	}

	// Create Temporal client
	client, err := temporal.NewClient(temporal.ClientOptions{
		HostPort:  config.Temporal.HostPort,
		Namespace: config.Temporal.Namespace,
		Identity:  config.Temporal.Identity,
		Logger:    logger,
	})
	if err != nil {
		logger.Fatal("Failed to create Temporal client", zap.Error(err))
	}
	defer client.Close()

	// Example of starting a simple workflow
	_, err = client.ExecuteWorkflow(temporal.StartWorkflowOptions{
		ID:        "simple-workflow",
		TaskQueue: "simple-task-queue",
	}, temporal.SimpleWorkflow)
	if err != nil {
		logger.Fatal("Failed to start simple workflow", zap.Error(err))
	}

	logger.Info("Simple workflow started successfully")

	// Example of starting a complex workflow
	_, err = client.ExecuteWorkflow(temporal.StartWorkflowOptions{
		ID:        "complex-workflow",
		TaskQueue: "complex-task-queue",
	}, temporal.ComplexWorkflow)
	if err != nil {
		logger.Fatal("Failed to start complex workflow", zap.Error(err))
	}

	logger.Info("Complex workflow started successfully")
}
