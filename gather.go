package main

import (
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/docker/go-dockercloud/dockercloud"
	c "github.com/infinityworks/docker-cloud-exporter/config"
	m "github.com/infinityworks/docker-cloud-exporter/measure"
	"github.com/prometheus/client_golang/prometheus"
)

func (e *Exporter) gatherStackMetrics(ch chan<- prometheus.Metric) (*dockercloud.Stack, error) {

	defer m.RecordFunction(time.Now(), "main", "gatherStackMetrics")

	var Stack = new(dockercloud.Stack)

	stack, err := dockercloud.ListStacks()
	if err != nil {
		log.Println(err)
	}

	log.Debugf("Data Captured", Stack)

	for _, x := range stack.Objects {

		e.gaugeVecs["stacksState"].With(prometheus.Labels{"stack_name": x.Name, "state": x.State}).Set(1)
		e.gaugeVecs["stackDeployedDate"].With(prometheus.Labels{"stack_name": x.Name}).Set(c.ConvertTime(x.Deployed_datetime))
	}

	return Stack, err

}

func (e *Exporter) gatherServiceMetrics(ch chan<- prometheus.Metric) (*dockercloud.Service, error) {

	defer m.RecordFunction(time.Now(), "main", "gatherServiceMetrics")

	var Service = new(dockercloud.Service)

	service, err := dockercloud.ListServices()
	if err != nil {
		log.Println(err)
	}

	log.Debugf("Data Captured", service)

	stack, err := dockercloud.ListStacks()
	if err != nil {
		log.Println(err)
	}

	for _, x := range service.Objects {

		var stackName = ""
		for _, y := range stack.Objects {
			if y.Resource_uri == x.Stack {
				stackName = y.Name
			}
		}

		e.gaugeVecs["serviceState"].With(prometheus.Labels{"service_name": x.Name, "stack_name": stackName, "state": x.State}).Set(1)
		e.gaugeVecs["serviceDeployedDate"].With(prometheus.Labels{"service_name": x.Name, "stack_name": stackName}).Set(c.ConvertTime(x.Deployed_datetime))
		e.gaugeVecs["serviceContainersRunning"].With(prometheus.Labels{"service_name": x.Name, "stack_name": stackName}).Set(float64(x.Current_num_containers))
		e.gaugeVecs["serviceContainersTarget"].With(prometheus.Labels{"service_name": x.Name, "stack_name": stackName}).Set(float64(x.Target_num_containers))
		e.gaugeVecs["serviceContainersStopped"].With(prometheus.Labels{"service_name": x.Name, "stack_name": stackName}).Set(float64(x.Stopped_num_containers))
	}

	return Service, err

}

func (e *Exporter) gatherNodeMetrics(ch chan<- prometheus.Metric) (*dockercloud.Node, error) {

	defer m.RecordFunction(time.Now(), "main", "gatherNodeMetrics")

	var Node = new(dockercloud.Node)

	node, err := dockercloud.ListNodes()
	if err != nil {
		log.Println(err)
	}

	log.Debugf("Data Captured", node)

	nodeCluster, err := dockercloud.ListNodeClusters()
	if err != nil {
		log.Println(err)
	}

	for _, x := range node.Objects {

		var nodeClusterName = ""
		for _, y := range nodeCluster.Objects {
			if y.Resource_uri == x.Node_cluster {
				nodeClusterName = y.Name
			}
		}

		e.gaugeVecs["nodeState"].With(prometheus.Labels{"node_uuid": x.Uuid, "availability_zone": x.Availability_zone, "region": x.Region, "node_cluster_name": nodeClusterName, "docker_version": x.Docker_version, "state": x.State}).Set(1)
	}

	return Node, err

}

func (e *Exporter) gatherNodeClusterMetrics(ch chan<- prometheus.Metric) (*dockercloud.NodeCluster, error) {

	defer m.RecordFunction(time.Now(), "main", "gatherNodeClusterMetrics")

	var NodeCluster = new(dockercloud.NodeCluster)

	nodeCluster, err := dockercloud.ListNodeClusters()
	if err != nil {
		log.Println(err)
	}

	log.Debugf("Data Captured", nodeCluster)

	for _, x := range nodeCluster.Objects {
		e.gaugeVecs["nodeClusterState"].With(prometheus.Labels{"node_cluster_name": x.Name, "state": x.State}).Set(1)
		e.gaugeVecs["nodeClusterNodesCurrent"].With(prometheus.Labels{"node_cluster_name": x.Name}).Set(float64(x.Current_num_nodes))
		e.gaugeVecs["nodeClusterNodesTarget"].With(prometheus.Labels{"node_cluster_name": x.Name}).Set(float64(x.Target_num_nodes))
	}

	return NodeCluster, err

}
