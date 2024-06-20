package temporal

// StartWorkflowOptions represents the options for starting a workflow
type StartWorkflowOptions struct {
	ID        string
	TaskQueue string
	// Additional options can be added as needed
}
