package temporal

import (
    "temporal-go-lib/internal/monitoring"
)

// ServeMetrics starts the HTTP server to expose Prometheus metrics
func ServeMetrics() {
    go monitoring.ServeMetrics()
}
