package main

import "github.com/prometheus/client_golang/prometheus"

func AddMetrics() map[string]*prometheus.GaugeVec {

	gaugeVecs := make(map[string]*prometheus.GaugeVec)

	// stack metrics
	gaugeVecs["stacksState"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "stack_state",
			Help:      "State of defined stack as reported by the Docker Cloud API",
		}, []string{"stack_name", "state"})
	gaugeVecs["stackDeployedDate"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "stack_deployed_date",
			Help:      "Date & time the stack was deployed in Unix seconds",
		}, []string{"stack_name"})

	// service metrics
	gaugeVecs["serviceState"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "service_state",
			Help:      "State of defined service as reported by the Docker Cloud API",
		}, []string{"service_name", "stack_name", "state"})
	gaugeVecs["serviceDeployedDate"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "service_deployed_date",
			Help:      "Date & time the service was created in Unix seconds",
		}, []string{"service_name", "stack_name"})
	gaugeVecs["serviceContainersRunning"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "service_containers_running",
			Help:      "Number of running containers",
		}, []string{"service_name", "stack_name"})
	gaugeVecs["serviceContainersTarget"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "service_containers_target",
			Help:      "Target number of containers defined in config",
		}, []string{"service_name", "stack_name"})
	gaugeVecs["serviceContainersStopped"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "service_containers_stopped",
			Help:      "Number of stopped containers",
		}, []string{"service_name", "stack_name"})

	// node metrics
	gaugeVecs["nodeState"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "node_state",
			Help:      "State of defined node as reported by the Docker Cloud API",
		}, []string{"node_uuid", "availability_zone", "region", "node_cluster_name", "docker_version", "state"})

	// node cluster metrics
	gaugeVecs["nodeClusterState"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "node_cluster_state",
			Help:      "State of defined node cluster as reported by the Docker Cloud API",
		}, []string{"node_cluster_name", "state"})
	gaugeVecs["nodeClusterNodesCurrent"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "node_cluster_nodes_current",
			Help:      "State of defined node cluster as reported by the Docker Cloud API",
		}, []string{"node_cluster_name"})
	gaugeVecs["nodeClusterNodesTarget"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "node_cluster_nodes_target",
			Help:      "State of defined node cluster as reported by the Docker Cloud API",
		}, []string{"node_cluster_name"})

	return gaugeVecs

}
