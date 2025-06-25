# Cloud-Native Go CLI & Service

A Go-based microservice with a CLI, Docker containerization, Kubernetes manifests, Helm chart, CI/CD pipeline, and Grafana dashboards.

## Features

- **Service**  
  - HTTP server with `/healthz` and `/hello` endpoints  
  - Prometheus metrics exposed at `/metrics`  
- **CLI** (using Cobra)  
  - `cloudcli status` — show current port  
  - `cloudcli purge` — stub for purging data  
- **Containerization**  
  - Multi-stage `Dockerfile` producing minimal binaries  
- **Orchestration**  
  - Kubernetes `Deployment` & `Service` manifests  
  - Ingress & TLS templating in Helm  
  - Horizontal Pod Autoscaler  
- **Packaging**  
  - Helm chart under `chart/cloud-service/`  
- **CI/CD**  
  - GitHub Actions workflow for build, test, Docker push, and Helm deploy  
- **Observability**  
  - Grafana dashboard auto-loaded from ConfigMap  

## Getting Started

1. Clone the repo  
   ```bash
   git clone https://github.com/venkatesh-gg/cloud-native-go-cli.git
   cd cloud-native-go-cli
