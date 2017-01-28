package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// FunctionDurations - Create a summary to track elapsed time of our key functions
	FunctionDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "function_durations_seconds",
			Help:       "Function timings for the Docker Cloud Exporter",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"pkg", "fnc"})
)

// Init registers the prometheus metrics for the measurement of the exporter itsself.
func Init() {
	prometheus.MustRegister(FunctionDurations)
}

func RecordFunction(start time.Time, pkgName string, fncName string) {
	elapsed := float64(time.Since(start))
	FunctionDurations.WithLabelValues(pkgName, fncName).Observe(elapsed)
}
