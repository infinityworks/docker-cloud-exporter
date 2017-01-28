package main

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// Exporter Sets up all the runtime and metrics
type Exporter struct {
	mutex     sync.RWMutex
	gaugeVecs map[string]*prometheus.GaugeVec
}

// NewExporter creates the metrics we wish to monitor
func newExporter() *Exporter {

	gaugeVecs := AddMetrics()

	return &Exporter{
		gaugeVecs: gaugeVecs,
	}
}
