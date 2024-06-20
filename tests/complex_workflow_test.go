package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap"
	"temporal-go-lib/pkg/temporal"
	"testing"
)

func TestComplexWorkflowFunctionality(t *testing.T) {
	// Initialize logger
	logger, err := zap.NewDevelopment()
	assert.NoError(t, err)
	defer logger.Sync()

	// Create Temporal client
	c, err := temporal.NewClient(temporal.ClientOptions{
		HostPort:  "localhost:7233",
		Namespace: "default",
		Logger:    logger,
	})
	assert.NoError(t, err)
	defer c.Close()

	// Create and start worker
	workerOptions := temporal.WorkerOptions{
		TaskQueue: "complex-task-queue",
		Logger:    logger,
	}
	w := temporal.NewWorker(c, workerOptions)
	w.RegisterWorkflow(temporal.ComplexWorkflow)
	w.RegisterActivity(temporal.ComplexActivity1)
	w.RegisterActivity(temporal.ComplexActivity2)
	assert.NoError(t, w.Start())
	defer w.Stop()

	// Execute workflow
	we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        "complex-workflow",
		TaskQueue: "complex-task-queue",
	}, temporal.ComplexWorkflow)
	assert.NoError(t, err)

	// Verify workflow result
	var result string
	err = we.Get(context.Background(), &result)
	assert.NoError(t, err)
	assert.Equal(t, "ComplexActivity1 completed for Step 1; ComplexActivity2 completed for Step 2", result)
}
