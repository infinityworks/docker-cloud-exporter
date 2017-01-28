# Example Metrics

Example of the metrics you could expect to see, returned for the service,stack, node and node cluster states.

```
# HELP docker_cloud_node_cluster_state State of defined node cluster as reported by the Docker Cloud API
# TYPE docker_cloud_node_cluster_state gauge
docker_cloud_node_cluster_state{node_cluster_name="testing",state="Deployed"} 1
# HELP docker_cloud_node_state State of defined node as reported by the Docker Cloud API
# TYPE docker_cloud_node_state gauge
docker_cloud_node_state{node_uuid="3e00675b-4dc0-4ea9-8329-aa852f34dcc7",state="Deployed"} 1
docker_cloud_node_state{node_uuid="c7f72de9-48ec-4db0-ae70-813c1ce2b7cf",state="Deployed"} 1
# HELP docker_cloud_service_created_date Date & time the service was created in Unix seconds
# TYPE docker_cloud_service_created_date gauge
docker_cloud_service_created_date{service_name="Example1",state="Running"} 1.485188537e+09
docker_cloud_service_created_date{service_name="grafana",state="Running"} 1.485188529e+09
docker_cloud_service_created_date{service_name="node-exporter",state="Running"} 1.485532594e+09
docker_cloud_service_created_date{service_name="prometheus",state="Running"} 1.485188566e+09
# HELP docker_cloud_service_state State of defined service as reported by the Docker Cloud API
# TYPE docker_cloud_service_state gauge
docker_cloud_service_state{service_name="Example1",state="Running"} 1
docker_cloud_service_state{service_name="grafana",state="Running"} 1
docker_cloud_service_state{service_name="node-exporter",state="Running"} 1
docker_cloud_service_state{service_name="prometheus",state="Running"} 1
# HELP docker_cloud_stack_created_date Date & time the stack was created in Unix seconds
# TYPE docker_cloud_stack_created_date gauge
docker_cloud_stack_created_date{stack_name="Example1",state="Running"} 1.485188688e+09
# HELP docker_cloud_stack_state State of defined stack as reported by the Docker Cloud API
# TYPE docker_cloud_stack_state gauge
docker_cloud_stack_state{stack_name="Example1",state="Running"} 1
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