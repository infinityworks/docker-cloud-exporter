package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
)

// Resets the guageVecs back to 0
// Ensures we start from a clean sheet
func (e *Exporter) resetMetrics() {

	for _, m := range e.gaugeVecs {
		m.Reset()
	}
}

// Describe describes all the metrics ever exported by the Docker Cloud exporter
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	for _, m := range e.gaugeVecs {
		m.Describe(ch)
	}
}

// Collect function, called on by Prometheus Client library
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	log.Info("Scrape recieved, collecting Metrics")

	e.mutex.Lock() // To protect metrics from concurrent collects.
	defer e.mutex.Unlock()

	e.resetMetrics() // Clean starting point

	e.gatherStackMetrics(ch)
	e.gatherServiceMetrics(ch)
	e.gatherNodeMetrics(ch)
	e.gatherNodeClusterMetrics(ch)

	for _, m := range e.gaugeVecs {
		m.Collect(ch)
	}

	log.Info("Metrics returned")

}
