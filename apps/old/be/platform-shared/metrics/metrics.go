package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Registry is the global Prometheus registry
	Registry = prometheus.NewRegistry()
)

// Counter creates a new Prometheus counter
func Counter(namespace, subsystem, name, help string, labels []string) *prometheus.CounterVec {
	return promauto.With(Registry).NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
		labels,
	)
}

// Histogram creates a new Prometheus histogram
func Histogram(namespace, subsystem, name, help string, buckets []float64, labels []string) *prometheus.HistogramVec {
	return promauto.With(Registry).NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
			Buckets:   buckets,
		},
		labels,
	)
}

// Gauge creates a new Prometheus gauge
func Gauge(namespace, subsystem, name, help string, labels []string) *prometheus.GaugeVec {
	return promauto.With(Registry).NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      name,
			Help:      help,
		},
		labels,
	)
}