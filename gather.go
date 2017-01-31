package main

import (
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/docker/go-dockercloud/dockercloud"
	c "github.com/infinityworksltd/docker-cloud-exporter/config"
	m "github.com/infinityworksltd/docker-cloud-exporter/measure"
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
		e.gaugeVecs["stackCreatedDate"].With(prometheus.Labels{"stack_name": x.Name, "state": x.State}).Set(c.ConvertTime(x.Deployed_datetime))
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

	for _, x := range service.Objects {

		e.gaugeVecs["serviceState"].With(prometheus.Labels{"service_name": x.Name, "stack_name": x.Stack, "state": x.State}).Set(1)
		e.gaugeVecs["serviceCreatedDate"].With(prometheus.Labels{"service_name": x.Name, "stack_name": x.Stack, "state": x.State}).Set(c.ConvertTime(x.Deployed_datetime))

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

	for _, x := range node.Objects {
		e.gaugeVecs["nodeState"].With(prometheus.Labels{"node_uuid": x.Uuid, "state": x.State}).Set(1)
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
	}

	return NodeCluster, err

}
