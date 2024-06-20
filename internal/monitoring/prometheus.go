package monitoring

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "net/http"
)

var (
    ClientConnections = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "temporal_client_connections",
        Help: "Number of active client connections",
    })
    RequestLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
        Name:    "temporal_client_request_latency_seconds",
        Help:    "Latency of client requests",
        Buckets: prometheus.DefBuckets,
    })
    ClientErrors = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "temporal_client_errors_total",
        Help: "Total number of client errors",
    })
    WorkerStart = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "temporal_worker_start_total",
        Help: "Total number of worker starts",
    })
    WorkerStop = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "temporal_worker_stop_total",
        Help: "Total number of worker stops",
    })
    WorkflowSuccess = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "temporal_workflow_success_total",
        Help: "Total number of successful workflows",
    })
    WorkflowErrors = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "temporal_workflow_errors_total",
        Help: "Total number of workflow errors",
    })
    ActivitySuccess = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "temporal_activity_success_total",
        Help: "Total number of successful activities",
    })
)

func init() {
    prometheus.MustRegister(ClientConnections, RequestLatency, ClientErrors, WorkerStart, WorkerStop, WorkflowSuccess, WorkflowErrors, ActivitySuccess)
}

func ServeMetrics() {
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":9090", nil)
}
