package main

import (
	"fmt"
	"go.uber.org/zap"
	"log"
	"temporal-go-lib/internal/config"
	"temporal-go-lib/internal/monitoring"
	"temporal-go-lib/pkg/temporal"
)

func main() {
	// Load configuration
	loadConfig, err := config.LoadConfig("loadConfig.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	var logger *zap.Logger
	if loadConfig.Logging.Level == "development" {
		logger, err = temporal.NewDevelopmentLogger()
	} else {
		logger, err = temporal.NewLogger()
	}
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	// Start monitoring if enabled
	if loadConfig.Monitoring.Enabled {
		go monitoring.ServeMetrics()
		logger.Info("Monitoring server started", zap.Int("port", loadConfig.Monitoring.Port))
	}

	// Create Temporal client
	client, err := temporal.NewClient(temporal.ClientOptions{
		HostPort:  loadConfig.Temporal.HostPort,
		Namespace: loadConfig.Temporal.Namespace,
		Identity:  loadConfig.Temporal.Identity,
		Logger:    logger,
	})
	if err != nil {
		logger.Fatal("Failed to create Temporal client", zap.Error(err))
	}
	defer client.Close()

	// Example of starting a simple workflow
	simpleRun, err := client.ExecuteWorkflow(temporal.StartWorkflowOptions{
		ID:        "simple-workflow",
		TaskQueue: "simple-task-queue",
	}, temporal.SimpleWorkflow)
	if err != nil {
		logger.Fatal("Failed to start simple workflow", zap.Error(err))
	}

	logger.Info(fmt.Sprintf("Simple workflow started successfully with Id: %v", simpleRun.GetRunID()))

	// Example of starting a complex workflow
	complexRun, err := client.ExecuteWorkflow(temporal.StartWorkflowOptions{
		ID:        "complex-workflow",
		TaskQueue: "complex-task-queue",
	}, temporal.ComplexWorkflow)
	if err != nil {
		logger.Fatal("Failed to start complex workflow", zap.Error(err))
	}

	logger.Info(fmt.Sprintf("Complex workflow started successfully with Id: %v", complexRun.GetRunID()))
}
