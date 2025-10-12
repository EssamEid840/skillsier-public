package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// HTTPMetrics holds HTTP RED metrics (Rate, Errors, Duration)
type HTTPMetrics struct {
	requestsTotal   *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	requestSize     *prometheus.HistogramVec
	responseSize    *prometheus.HistogramVec
}

// NewHTTPMetrics creates HTTP metrics
func NewHTTPMetrics(namespace, subsystem string) *HTTPMetrics {
	return &HTTPMetrics{
		requestsTotal: Counter(
			namespace, subsystem, "http_requests_total",
			"Total number of HTTP requests",
			[]string{"method", "path", "status"},
		),
		requestDuration: Histogram(
			namespace, subsystem, "http_request_duration_seconds",
			"HTTP request duration in seconds",
			[]float64{.001, .005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
			[]string{"method", "path", "status"},
		),
		requestSize: Histogram(
			namespace, subsystem, "http_request_size_bytes",
			"HTTP request size in bytes",
			prometheus.ExponentialBuckets(100, 10, 8),
			[]string{"method", "path"},
		),
		responseSize: Histogram(
			namespace, subsystem, "http_response_size_bytes",
			"HTTP response size in bytes",
			prometheus.ExponentialBuckets(100, 10, 8),
			[]string{"method", "path", "status"},
		),
	}
}

// Middleware returns Gin middleware for HTTP metrics
func (m *HTTPMetrics) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		// Request size
		if c.Request.ContentLength > 0 {
			m.requestSize.WithLabelValues(c.Request.Method, path).Observe(float64(c.Request.ContentLength))
		}

		c.Next()

		// After request
		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())

		m.requestsTotal.WithLabelValues(c.Request.Method, path, status).Inc()
		m.requestDuration.WithLabelValues(c.Request.Method, path, status).Observe(duration)
		m.responseSize.WithLabelValues(c.Request.Method, path, status).Observe(float64(c.Writer.Size()))
	}
}