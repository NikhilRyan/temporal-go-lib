# Temporal Go Library

This library provides a user-friendly interface for using Temporal in Go. It supports all types of workflows, activities, and provides a modular structure for easy integration. The library includes advanced features such as structured logging, monitoring, error handling, security, and configuration management.

## Installation

```bash
go get github.com/NikhilRyan/temporal-go-lib
```

## Usage

### Configuration

Create a `config.yaml` file with the following structure:

```yaml
temporal:
  hostPort: "localhost:7233"
  namespace: "default"
  identity: "your-identity"

retryPolicy:
  initialInterval: 1
  backoffCoefficient: 2.0
  maximumInterval: 60
  maximumAttempts: 5

logging:
  level: "production" # or "development"

monitoring:
  enabled: true
  port: 9090
```

### Creating a Temporal Client

```go
import (
    "github.com/NikhilRyan/temporal-go-lib/pkg/temporal"
    "log"
)

func main() {
    client, err := temporal.NewClient(temporal.ClientOptions{
        HostPort: "localhost:7233",
        Namespace: "default",
    })
    if err != nil {
        log.Fatalf("Failed to create Temporal client: %v", err)
    }
    defer client.Close()
}
```

### Starting a Simple Workflow

```go
err := client.ExecuteWorkflow(temporal.StartWorkflowOptions{
    ID:        "simple-workflow",
    TaskQueue: "simple-task-queue",
}, temporal.SimpleWorkflow)
if err != nil {
    log.Fatalf("Failed to start workflow: %v", err)
}
```

### Running a Worker

```go
workerOptions := temporal.WorkerOptions{
    TaskQueue: "simple-task-queue",
    MaxConcurrentActivityExecutionSize: 10,
    MaxConcurrentWorkflowTaskExecutionSize: 10,
    WorkerActivitiesPerSecond: 100.0,
    WorkerStopTimeout: time.Minute,
    Logger: logger,
}

worker := temporal.NewWorker(client, workerOptions)
worker.RegisterWorkflow(temporal.SimpleWorkflow)
worker.RegisterActivity(temporal.SimpleActivity)

err = worker.Start()
if err != nil {
    logger.Fatal("Failed to start Temporal worker", zap.Error(err))
}
defer worker.Stop()

logger.Info("Worker started successfully")
```

### Examples

See the `examples/` directory for more examples.

- [Simple Workflow Example](examples/simple_workflow/main.go)
- [Base Workflow Example](examples/base_workflow/main.go)
- [Complex Workflow Example](examples/complex_workflow/main.go)

### Directory Structure

```
temporal-go-lib/
├── README.md
├── go.mod
├── go.sum
├── main.go
├── config.yaml
├── examples/
│   ├── base_workflow/
│   │   └── main.go
│   ├── complex_workflow/
│   │   └── main.go
│   └── simple_workflow/
│       └── main.go
├── internal/
│   ├── client/
│   │   └── client.go
│   ├── worker/
│   │   └── worker.go
│   ├── workflows/
│   │   ├── base_workflow.go
│   │   ├── complex_workflow.go
│   │   └── simple_workflow.go
│   ├── activities/
│   │   └── activity.go
│   ├── logging/
│   │   └── logging.go
│   ├── monitoring/
│   │   └── prometheus.go
│   ├── config/
│   │   └── config.go
│   ├── errorhandling/
│   │   └── errors.go
│   └── security/
│       └── security.go
├── pkg/
│   └── temporal/
│       ├── client.go
│       ├── worker.go
│       ├── workflow.go
│       ├── activity.go
│       ├── logging.go
│       ├── monitoring.go
│       ├── errorhandling.go
│       └── security.go
└── vendor/
```

## Features

### Logging

Structured logging using `zap` for better traceability and insights.

### Monitoring

Integration with Prometheus for monitoring and metrics collection.

### Error Handling

Enhanced error handling and retry mechanisms.

### Configuration Management

Flexible configuration management using a `config.yaml` file.

### Security

Features like encrypted payloads and secure communication.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss your changes.

## License

This project is licensed under the MIT License.
