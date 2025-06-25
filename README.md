# Cloud-Native Go CLI & Service

A Go-based microservice with a CLI, Docker containerization, Kubernetes manifests, Helm chart, CI/CD pipeline, and Grafana dashboards.

## Features

* **Service**

  * HTTP server with `/healthz` and `/hello` endpoints
  * Prometheus metrics exposed at `/metrics`
* **CLI** (using Cobra)

  * `cloudcli status` — show current port
  * `cloudcli purge` — stub for purging data
* **Containerization**

  * Multi-stage `Dockerfile` producing minimal binaries
* **Orchestration**

  * Kubernetes `Deployment` & `Service` manifests
  * Ingress & TLS templating in Helm
  * Horizontal Pod Autoscaler
* **Packaging**

  * Helm chart under `chart/cloud-service/`
* **CI/CD**

  * GitHub Actions workflow for build, test, Docker push, and Helm deploy
* **Observability**

  * Grafana dashboard auto-loaded from ConfigMap

## Getting Started

1. **Clone the repo**

   ```bash
   git clone https://github.com/venkatesh-gg/cloud-native-go-cli.git
   cd cloud-native-go-cli
   ```

2. **Run the service locally**

   ```bash
   go run cmd/server/main.go
   curl http://localhost:8080/healthz
   ```

3. **Build & run the CLI**

   ```bash
   go build -o cloudcli ./cmd/cli
   ./cloudcli status
   ```

4. **Build the Docker image**

   ```bash
   docker build -t venkatesg559/cloud-service:latest .
   ```

5. **Kubernetes (Docker Desktop or Minikube)**

   ```bash
   kubectl apply -f k8s/deployment.yaml
   kubectl apply -f k8s/service.yaml
   kubectl port-forward svc/cloud-service 9090:80
   curl http://localhost:9090/healthz
   ```

6. **Helm install/upgrade**

   ```bash
   helm upgrade --install cloud-service chart/cloud-service \
     --set image.repository=venkatesg559/cloud-service \
     --set image.tag=latest
   ```

## CI/CD

On each push to `main`, GitHub Actions will:

1. Run Go tests
2. Build and push Docker image
3. Upgrade Helm release

## Grafana Dashboard

1. Install Grafana in the `monitoring` namespace
2. Enable the dashboards sidecar
3. Apply `k8s/grafana-dashboard.yaml`
4. View at [http://localhost:3000](http://localhost:3000) (admin/admin)

---

## Contributing

Feel free to open issues or submit pull requests — happy to collaborate!
