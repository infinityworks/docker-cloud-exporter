package main

import "github.com/prometheus/client_golang/prometheus"

func AddMetrics() map[string]*prometheus.GaugeVec {

	gaugeVecs := make(map[string]*prometheus.GaugeVec)

	// stack metrics
	gaugeVecs["stacksState"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "docker_cloud",
			Name:      "stack_state",
			Help:      "State of defined stack as reported by the Docker Cloud API",
		}, []string{"stack_name", "state"})
	gaugeVecs["stackCreatedDate"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "docker_cloud",
			Name:      "stack_created_date",
			Help:      "Date & time the stack was created in Unix seconds",
		}, []string{"stack_name", "state"})

	// service metrics
	gaugeVecs["serviceState"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "docker_cloud",
			Name:      "service_state",
			Help:      "State of defined service as reported by the Docker Cloud API",
		}, []string{"service_name", "stack_name", "state"})
	gaugeVecs["serviceCreatedDate"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "docker_cloud",
			Name:      "service_created_date",
			Help:      "Date & time the service was created in Unix seconds",
		}, []string{"service_name", "stack_name". "state"})

	// node metrics
	gaugeVecs["nodeState"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "docker_cloud",
			Name:      "node_state",
			Help:      "State of defined node as reported by the Docker Cloud API",
		}, []string{"node_uuid", "state"})

	// node cluster metrics
	gaugeVecs["nodeClusterState"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "docker_cloud",
			Name:      "node_cluster_state",
			Help:      "State of defined node cluster as reported by the Docker Cloud API",
		}, []string{"node_cluster_name", "state"})

	return gaugeVecs

}
