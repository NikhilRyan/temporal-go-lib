package client

import (
    "go.temporal.io/sdk/client"
    "go.uber.org/zap"
    "time"
    "temporal-go-lib/internal/monitoring"
)

// ClientOptions holds configuration for the Temporal client
type ClientOptions struct {
    HostPort      string
    Namespace     string
    RetryPolicy   RetryPolicy
    MetricsScope  string
    Identity      string
    DataConverter client.DataConverter
    Logger        *zap.Logger
}

// RetryPolicy defines the retry options for Temporal client
type RetryPolicy struct {
    InitialInterval    time.Duration
    BackoffCoefficient float64
    MaximumInterval    time.Duration
    MaximumAttempts    int
}

// Client wraps the Temporal SDK client
type Client struct {
    temporalClient client.Client
    logger         *zap.Logger
}

// NewClient creates a new Temporal client
func NewClient(options ClientOptions) (*Client, error) {
    c, err := client.Dial(client.Options{
        HostPort:      options.HostPort,
        Namespace:     options.Namespace,
        MetricsScope:  nil, // Placeholder for actual metrics scope
        Identity:      options.Identity,
        DataConverter: options.DataConverter,
        RetryPolicy: &client.RetryPolicy{
            InitialInterval:    options.RetryPolicy.InitialInterval,
            BackoffCoefficient: options.RetryPolicy.BackoffCoefficient,
            MaximumInterval:    options.RetryPolicy.MaximumInterval,
            MaximumAttempts:    options.RetryPolicy.MaximumAttempts,
        },
    })
    if err != nil {
        options.Logger.Error("Failed to create Temporal client", zap.Error(err))
        monitoring.ClientErrors.Inc()
        return nil, err
    }
    options.Logger.Info("Temporal client created successfully")
    monitoring.ClientConnections.Inc()
    return &Client{temporalClient: c, logger: options.Logger}, nil
}

// Close closes the Temporal client
func (c *Client) Close() {
    c.temporalClient.Close()
    c.logger.Info("Temporal client closed successfully")
    monitoring.ClientConnections.Dec()
}

// ExecuteWorkflow executes a workflow
func (c *Client) ExecuteWorkflow(options client.StartWorkflowOptions, workflow interface{}, args ...interface{}) (client.WorkflowRun, error) {
    c.logger.Info("Executing workflow", zap.String("WorkflowID", options.ID))
    return c.temporalClient.ExecuteWorkflow(client.BackgroundContext(), options, workflow, args...)
}

// SignalWithStartWorkflow sends a signal to a workflow or starts it if it does not exist
func (c *Client) SignalWithStartWorkflow(options client.StartWorkflowOptions, signalName string, signalArg interface{}, workflow interface{}, args ...interface{}) (client.WorkflowRun, error) {
    c.logger.Info("SignalWithStartWorkflow", zap.String("WorkflowID", options.ID), zap.String("SignalName", signalName))
    return c.temporalClient.SignalWithStartWorkflow(client.BackgroundContext(), options, signalName, signalArg, workflow, args...)
}

// QueryWorkflow queries a workflow
func (c *Client) QueryWorkflow(workflowID, runID, queryType string, args ...interface{}) (client.WorkflowQueryResult, error) {
    c.logger.Info("Querying workflow", zap.String("WorkflowID", workflowID), zap.String("RunID", runID), zap.String("QueryType", queryType))
    return c.temporalClient.QueryWorkflow(client.BackgroundContext(), workflowID, runID, queryType, args...)
}

// GetWorkflowHistory gets the history of a workflow
func (c *Client) GetWorkflowHistory(workflowID, runID string) ([]*client.HistoryEvent, error) {
    iterator := c.temporalClient.GetWorkflowHistory(client.BackgroundContext(), workflowID, runID, false, client.HistoryEventFilterTypeAllEvent)
    var history []*client.HistoryEvent
    for iterator.HasNext() {
        event, err := iterator.Next()
        if err != nil {
            return nil, err
        }
        history = append(history, event)
    }
    return history, nil
}
