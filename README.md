# docker-cloud-exporter

Exposes the health of Stacks, Services, Nodes and Node Clusters from the Docker Cloud API, to a Prometheus compatible endpoint.


## Description

The application can be run in a number of ways, the main consumption is the Docker hub image `infinityworksltd/docker-cloud-exporter`. 

**Required**
* `DOCKERCLOUD_USER`      // Username of your docker cloud account
* `DOCKERCLOUD_APIKEY`    // Generated through the docker cloud admin interface

**Optional**
* `DOCKERCLOUD_NAMESPACE` // Specify this if you wish to show metrics for an organisation your user has access to. Specify the name of the org in this env var.
* `METRICS_PATH`          // Path under which to expose metrics.
* `LISTEN_PORT`           // Port on which to expose metrics. Defaults to 9174
*	`LOG_LEVEL`           // Optional - Set the logging level, defaults to Info

## Install and deploy

Run manually from Docker Hub:
```
docker run -d -e DOCKERCLOUD_USER="XXXXXXXX" -e DOCKERCLOUD_APIKEY="XXXXXXX" -p 9174:9174 infinityworks/docker-cloud-exporter
```

Build a docker image:
```
docker build -t <image-name> .
docker run -d -e DOCKERCLOUD_USER="XXXXXXXX" -e DOCKERCLOUD_APIKEY="XXXXXXX" -p 9174:9174 <image-name>
```

## Docker compose

```
docker-cloud-exporter:
    tty: true
    stdin_open: true
    environment:
      - DOCKERCLOUD_USER="xxxx"
      - DOCKERCLOUD_APIKEY="xxxxxx"
    expose:
      - 9174:9174
    image: infinityworks/docker-cloud-exporter:latest
```


## Metrics

Metrics will be made available on port 9174 by default, or you can pass environment variable ```LISTEN_ADDRESS``` to override this.
An example printout of the metrics you should expect to see can be found in `METRICS.md`.


## Metadata
[![](https://images.microbadger.com/badges/version/infinityworks/docker-cloud-exporter.svg)](http://microbadger.com/images/infinityworks/docker-cloud-exporter "Get your own version badge on microbadger.com") [![](https://images.microbadger.com/badges/image/infinityworks/docker-cloud-exporter.svg)](http://microbadger.com/images/infinityworks/docker-cloud-exporter "Get your own image badge on microbadger.com")
[![Go Report Card](https://goreportcard.com/badge/github.com/infinityworksltd/docker-cloud-exporter)](https://goreportcard.com/report/github.com/infinityworksltd/docker-cloud-exporter)
[![GoDoc](https://godoc.org/github.com/infinityworksltd/docker-cloud-exporter?status.svg)](https://godoc.org/github.com/infinityworksltd/docker-cloud-exporter)