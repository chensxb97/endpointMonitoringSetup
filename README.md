# endpointMonitoringSetup
Exploring the different ways to set up blackbox exporter for endpoint monitoring.

## Running Blackbox Exporter
Blackbox exporter can be first installed by either installing the binary from [here](https://github.com/prometheus/blackbox_exporter/releases) or cloning the [repo](https://github.com/prometheus/blackbox_exporter.git). Afterwards, proceed with the following:

1. Modify module configuration -> `./local_test/blackbox.yml`.

2. Run Blackbox Exporter.
**Binary**
```shell
./blackbox_exporter --config.file=./local_test/blackbox.yml
```

**From Scratch**
```shell
cd blackbox_exporter/ && go run main.go --config.file=./local_test/blackbox.yml
```

## Running Prometheus
Similarly, you can choose to either install the binary from [here](https://prometheus.io/download/) or clone the [repo](https://github.com/prometheus/prometheus).

1. Define configuration files. This project explores the following three implementations, which are instrumented using three different configuration files.
    - `./local_test/static.yml` - Monitoring using a static list of targets
    - `./local_test/service_discovery.yml` - Monitoring a dynamic list of targets using HTTP Service Discovery
    - `./local_test/service_discovery_custom_labels.yml` - Monitoring a dynamic list of targets using HTTP Service Discovery + Introducing new custom labels.
2. Run Prometheus.

**Binary**
```shell
./prometheus --config.file=../local_test/static.yml
```

**From Scratch**
```shell
cd prometheus/ && make build
./prometheus --config.file=./local_test/static.yml
```

## Run HTTP Endpoint for Service Discovery
```go
// HTTP SD endpoint will be discoverable on localhost:8000/targets.
cd backend/ && go run main.go
```
