package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	c "github.com/infinityworks/docker-cloud-exporter/config"
	m "github.com/infinityworks/docker-cloud-exporter/measure"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "docker_cloud" // Used to prepand Prometheus metrics
)

// Runtime variables
var (
	metricsPath = c.GetEnv("METRICS_PATH", "/metrics") // Path under which to expose metrics
	listenPort  = c.GetEnv("LISTEN_PORT", ":9174")     // Port on which to expose metrics
	logLevel    = c.GetEnv("LOG_LEVEL", "info")        // Optional - Set the logging level
)

func main() {

	c.CheckConfig()

	setLogLevel(logLevel)
	log.Info("Starting Prometheus Exporter")

	// Register internal metrics used for tracking function performance
	m.Init()

	Exporter := newExporter()
	prometheus.MustRegister(Exporter)

	http.Handle(metricsPath, prometheus.Handler())

	log.Printf("Starting Server on port %s and path %s", listenPort, metricsPath)
	log.Fatal(http.ListenAndServe(listenPort, nil))

}
