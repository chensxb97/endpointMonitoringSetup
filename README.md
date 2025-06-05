# endpointMonitoringSetup
Exploring the different ways to set up Blackbox exporter/Prometheus for endpoint monitoring.

## Running Blackbox Exporter
Blackbox exporter can be first installed by either installing the binary from [here](https://github.com/prometheus/blackbox_exporter/releases) or cloning the [repo](https://github.com/prometheus/blackbox_exporter.git). Afterwards, proceed with the following:

1. Modify module configuration - `./configs/blackbox.yml`.

2. Run Blackbox Exporter.

**Running from Binary**
```shell
./blackbox_exporter --config.file=./configs/blackbox.yml
```

**Running from repo**
```shell
cd blackbox_exporter/ && go run main.go --config.file=./configs/blackbox.yml
```

## Running Prometheus
Similarly, you can choose to either install the binary from [here](https://prometheus.io/download/) or clone the [repo](https://github.com/prometheus/prometheus).

1. Define configuration files. This project explores the following three implementations, which are instrumented using three different configuration files.
    - `./configs/static.yml` - Defines static list of targets
    - `./configs/service_discovery.yml` - Defines a HTTP Service Discovery to fetch dynamic targets
    - `./configs/service_discovery_custom_labels.yml` - Monitoring HTTP Service Discovery to fetch dynamic targets with additional custom labels.

2. Run Prometheus.

**Running from Binary**
```shell
./prometheus --config.file=./configs/static.yml
```

**Build and run manually**
```shell
cd prometheus/ && make build
./prometheus --config.file=./configs/static.yml
```

## Run HTTP Endpoints
For this project, a simple backend script in Go is used to simulate HTTP SD endpoints. They are are discoverable on `localhost:8000/targets` and `localhost:8000/targets_custom`.
```go
cd backend/ && go run main.go
```
