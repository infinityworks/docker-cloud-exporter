# Example Metrics

Example of the metrics you could expect to see, returned for the service,stack, node and node cluster states.

```
# HELP docker_cloud_node_cluster_nodes_current State of defined node cluster as reported by the Docker Cloud API
# TYPE docker_cloud_node_cluster_nodes_current gauge
docker_cloud_node_cluster_nodes_current{node_cluster_name="testing"} 2
# HELP docker_cloud_node_cluster_nodes_target State of defined node cluster as reported by the Docker Cloud API
# TYPE docker_cloud_node_cluster_nodes_target gauge
docker_cloud_node_cluster_nodes_target{node_cluster_name="testing"} 2
# HELP docker_cloud_node_cluster_state State of defined node cluster as reported by the Docker Cloud API
# TYPE docker_cloud_node_cluster_state gauge
docker_cloud_node_cluster_state{node_cluster_name="testing",state="Deployed"} 1
# HELP docker_cloud_node_state State of defined node as reported by the Docker Cloud API
# TYPE docker_cloud_node_state gauge
docker_cloud_node_state{availability_zone="",docker_version="1.11.2-cs5",node_cluster_name="ExampleStack",node_uuid="sssssss-4dc0-4ea9-8329-aa852f34dcc7",region="/api/infra/v1/region/digitalocean/ams3/",state="Deployed"} 1
docker_cloud_node_state{availability_zone="",docker_version="1.11.2-cs5",node_cluster_name="ExampleStack",node_uuid="c7fssde9-48ec-4db0-ae70-813c1ce2b7cf",region="/api/infra/v1/region/digitalocean/ams3/",state="Deployed"} 1
# HELP docker_cloud_service_containers_running Number of running containers
# TYPE docker_cloud_service_containers_running gauge
docker_cloud_service_containers_running{service_name="example_service",stack_name="ExampleStack"} 1
docker_cloud_service_containers_running{service_name="docker-cloud-exporter",stack_name="ExampleStack"} 1
docker_cloud_service_containers_running{service_name="github-exporter",stack_name="ExampleStack"} 1
docker_cloud_service_containers_running{service_name="grafana",stack_name="ExampleStack"} 1
docker_cloud_service_containers_running{service_name="node-exporter",stack_name="ExampleStack"} 2
docker_cloud_service_containers_running{service_name="prometheus",stack_name="ExampleStack"} 1
# HELP docker_cloud_service_containers_stopped Number of stopped containers
# TYPE docker_cloud_service_containers_stopped gauge
docker_cloud_service_containers_stopped{service_name="example_service",stack_name="ExampleStack"} 0
docker_cloud_service_containers_stopped{service_name="docker-cloud-exporter",stack_name="ExampleStack"} 0
docker_cloud_service_containers_stopped{service_name="github-exporter",stack_name="ExampleStack"} 0
docker_cloud_service_containers_stopped{service_name="grafana",stack_name="ExampleStack"} 0
docker_cloud_service_containers_stopped{service_name="node-exporter",stack_name="ExampleStack"} 0
docker_cloud_service_containers_stopped{service_name="prometheus",stack_name="ExampleStack"} 0
# HELP docker_cloud_service_containers_target Target number of containers defined in config
# TYPE docker_cloud_service_containers_target gauge
docker_cloud_service_containers_target{service_name="example_service",stack_name="ExampleStack"} 1
docker_cloud_service_containers_target{service_name="docker-cloud-exporter",stack_name="ExampleStack"} 1
docker_cloud_service_containers_target{service_name="github-exporter",stack_name="ExampleStack"} 1
docker_cloud_service_containers_target{service_name="grafana",stack_name="ExampleStack"} 1
docker_cloud_service_containers_target{service_name="node-exporter",stack_name="ExampleStack"} 0
docker_cloud_service_containers_target{service_name="prometheus",stack_name="ExampleStack"} 1
# HELP docker_cloud_service_deployed_date Date & time the service was created in Unix seconds
# TYPE docker_cloud_service_deployed_date gauge
docker_cloud_service_deployed_date{service_name="example_service",stack_name="ExampleStack"} 1.485188537e+09
docker_cloud_service_deployed_date{service_name="docker-cloud-exporter",stack_name="ExampleStack"} 1.485618065e+09
docker_cloud_service_deployed_date{service_name="github-exporter",stack_name="ExampleStack"} 1.485726059e+09
docker_cloud_service_deployed_date{service_name="grafana",stack_name="ExampleStack"} 1.485188529e+09
docker_cloud_service_deployed_date{service_name="node-exporter",stack_name="ExampleStack"} 1.485532594e+09
docker_cloud_service_deployed_date{service_name="prometheus",stack_name="ExampleStack"} 1.485188566e+09
# HELP docker_cloud_service_state State of defined service as reported by the Docker Cloud API
# TYPE docker_cloud_service_state gauge
docker_cloud_service_state{service_name="example_service",stack_name="ExampleStack",state="Running"} 1
docker_cloud_service_state{service_name="docker-cloud-exporter",stack_name="ExampleStack",state="Running"} 1
docker_cloud_service_state{service_name="github-exporter",stack_name="ExampleStack",state="Running"} 1
docker_cloud_service_state{service_name="grafana",stack_name="ExampleStack",state="Running"} 1
docker_cloud_service_state{service_name="node-exporter",stack_name="ExampleStack",state="Running"} 1
docker_cloud_service_state{service_name="prometheus",stack_name="ExampleStack",state="Running"} 1
# HELP docker_cloud_stack_deployed_date Date & time the stack was deployed in Unix seconds
# TYPE docker_cloud_stack_deployed_date gauge
docker_cloud_stack_deployed_date{stack_name="ExampleStack"} 1.485188688e+09
# HELP docker_cloud_stack_state State of defined stack as reported by the Docker Cloud API
# TYPE docker_cloud_stack_state gauge
docker_cloud_stack_state{stack_name="ExampleStack",state="Running"} 1
```

An example of the internal metrics to track the performance of the exporter, and useful as a basic example how to instrument your code.

```
# HELP function_count_totals total count of function calls
# TYPE function_count_totals counter
function_count_totals{fnc="getJSON",pkg="hosts"} 3
function_count_totals{fnc="getJSON",pkg="services"} 3
function_count_totals{fnc="getJSON",pkg="stacks"} 3
# HELP function_durations_seconds Function timings for Rancher Exporter
# TYPE function_durations_seconds summary
function_durations_seconds{fnc="getJSON",pkg="hosts",quantile="0.5"} 33546
function_durations_seconds{fnc="getJSON",pkg="hosts",quantile="0.9"} 59199
function_durations_seconds{fnc="getJSON",pkg="hosts",quantile="0.99"} 59199
function_durations_seconds_sum{fnc="getJSON",pkg="hosts"} 121502
function_durations_seconds_count{fnc="getJSON",pkg="hosts"} 3
function_durations_seconds{fnc="getJSON",pkg="services",quantile="0.5"} 49354
function_durations_seconds{fnc="getJSON",pkg="services",quantile="0.9"} 63310
function_durations_seconds{fnc="getJSON",pkg="services",quantile="0.99"} 63310
function_durations_seconds_sum{fnc="getJSON",pkg="services"} 146519
function_durations_seconds_count{fnc="getJSON",pkg="services"} 3
function_durations_seconds{fnc="getJSON",pkg="stacks",quantile="0.5"} 45075
function_durations_seconds{fnc="getJSON",pkg="stacks",quantile="0.9"} 59805
function_durations_seconds{fnc="getJSON",pkg="stacks",quantile="0.99"} 59805
function_durations_seconds_sum{fnc="getJSON",pkg="stacks"} 134789
function_durations_seconds_count{fnc="getJSON",pkg="stacks"} 3
```