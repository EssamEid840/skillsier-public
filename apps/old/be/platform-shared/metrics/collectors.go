package metrics

import (
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
)

// DBStatsCollector collects database pool statistics
type DBStatsCollector struct {
	db *sql.DB

	maxOpenConnections *prometheus.Desc
	openConnections    *prometheus.Desc
	inUse              *prometheus.Desc
	idle               *prometheus.Desc
	waitCount          *prometheus.Desc
	waitDuration       *prometheus.Desc
	maxIdleClosed      *prometheus.Desc
	maxLifetimeClosed  *prometheus.Desc
}

// NewDBStatsCollector creates a new DB stats collector
func NewDBStatsCollector(namespace, subsystem string, db *sql.DB) *DBStatsCollector {
	return &DBStatsCollector{
		db: db,
		maxOpenConnections: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "db_max_open_connections"),
			"Maximum number of open connections to the database",
			nil, nil,
		),
		openConnections: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "db_open_connections"),
			"Number of established connections to the database",
			nil, nil,
		),
		inUse: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "db_in_use_connections"),
			"Number of connections currently in use",
			nil, nil,
		),
		idle: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "db_idle_connections"),
			"Number of idle connections",
			nil, nil,
		),
		waitCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "db_wait_count_total"),
			"Total number of connections waited for",
			nil, nil,
		),
		waitDuration: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "db_wait_duration_seconds_total"),
			"Total time blocked waiting for a new connection",
			nil, nil,
		),
		maxIdleClosed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "db_max_idle_closed_total"),
			"Total number of connections closed due to SetMaxIdleConns",
			nil, nil,
		),
		maxLifetimeClosed: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "db_max_lifetime_closed_total"),
			"Total number of connections closed due to SetConnMaxLifetime",
			nil, nil,
		),
	}
}

// Describe implements prometheus.Collector
func (c *DBStatsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.maxOpenConnections
	ch <- c.openConnections
	ch <- c.inUse
	ch <- c.idle
	ch <- c.waitCount
	ch <- c.waitDuration
	ch <- c.maxIdleClosed
	ch <- c.maxLifetimeClosed
}

// Collect implements prometheus.Collector
func (c *DBStatsCollector) Collect(ch chan<- prometheus.Metric) {
	stats := c.db.Stats()

	ch <- prometheus.MustNewConstMetric(c.maxOpenConnections, prometheus.GaugeValue, float64(stats.MaxOpenConnections))
	ch <- prometheus.MustNewConstMetric(c.openConnections, prometheus.GaugeValue, float64(stats.OpenConnections))
	ch <- prometheus.MustNewConstMetric(c.inUse, prometheus.GaugeValue, float64(stats.InUse))
	ch <- prometheus.MustNewConstMetric(c.idle, prometheus.GaugeValue, float64(stats.Idle))
	ch <- prometheus.MustNewConstMetric(c.waitCount, prometheus.CounterValue, float64(stats.WaitCount))
	ch <- prometheus.MustNewConstMetric(c.waitDuration, prometheus.CounterValue, stats.WaitDuration.Seconds())
	ch <- prometheus.MustNewConstMetric(c.maxIdleClosed, prometheus.CounterValue, float64(stats.MaxIdleClosed))
	ch <- prometheus.MustNewConstMetric(c.maxLifetimeClosed, prometheus.CounterValue, float64(stats.MaxLifetimeClosed))
}