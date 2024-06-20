package temporal

import (
	"context"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

// Client wraps the Temporal SDK client
type Client struct {
	internalClient client.Client
	logger         *zap.Logger
}

// ClientOptions holds configuration for the Temporal client
type ClientOptions struct {
	HostPort  string
	Namespace string
	Identity  string
	Logger    *zap.Logger
}

// NewClient creates a new Temporal client
func NewClient(options ClientOptions) (*Client, error) {
	internalClient, err := client.NewClient(client.Options{
		HostPort:  options.HostPort,
		Namespace: options.Namespace,
		Identity:  options.Identity,
	})
	if err != nil {
		options.Logger.Error("Failed to create Temporal client", zap.Error(err))
		return nil, err
	}
	options.Logger.Info("Temporal client created successfully")
	return &Client{internalClient: internalClient, logger: options.Logger}, nil
}

// Close closes the Temporal client
func (c *Client) Close() {
	c.internalClient.Close()
	c.logger.Info("Temporal client closed successfully")
}

// ExecuteWorkflow executes a workflow
func (c *Client) ExecuteWorkflow(options StartWorkflowOptions, workflow interface{}, args ...interface{}) (client.WorkflowRun, error) {
	c.logger.Info("Executing workflow", zap.String("WorkflowID", options.ID))
	return c.internalClient.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        options.ID,
		TaskQueue: options.TaskQueue,
	}, workflow, args...)
}

// SignalWithStartWorkflow sends a signal to a workflow or starts it if it does not exist
func (c *Client) SignalWithStartWorkflow(options StartWorkflowOptions, signalName string, signalArg interface{}, workflow interface{}, args ...interface{}) (client.WorkflowRun, error) {
	c.logger.Info("SignalWithStartWorkflow", zap.String("WorkflowID", options.ID), zap.String("SignalName", signalName))
	return c.internalClient.SignalWithStartWorkflow(context.Background(), client.StartWorkflowOptions{
		ID:        options.ID,
		TaskQueue: options.TaskQueue,
	}, signalName, signalArg, workflow, args...)
}

// QueryWorkflow queries a workflow
func (c *Client) QueryWorkflow(workflowID, runID, queryType string, args ...interface{}) (client.WorkflowRun, error) {
	c.logger.Info("Querying workflow", zap.String("WorkflowID", workflowID), zap.String("RunID", runID), zap.String("QueryType", queryType))
	return c.internalClient.QueryWorkflow(context.Background(), workflowID, runID, queryType, args...)
}

// GetWorkflowHistory gets the history of a workflow
func (c *Client) GetWorkflowHistory(workflowID, runID string) client.HistoryEventIterator {
	iterator := c.internalClient.GetWorkflowHistory(context.Background(), workflowID, runID, false, client.HistoryEventFilterTypeAllEvent)
	return iterator
}
