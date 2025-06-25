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
2. Run the service locally
   go run cmd/server/main.go
curl http://localhost:8080/healthz
3. Build & run the CLI
   go build -o cloudcli ./cmd/cli
./cloudcli status
4. Build the Docker image
   docker build -t venkatesg559/cloud-service:latest .
5. Kubernetes (Docker Desktop or Minikube)
   kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl port-forward svc/cloud-service 9090:80
curl http://localhost:9090/healthz
6. Helm install/upgrade
   helm upgrade --install cloud-service chart/cloud-service \
  --set image.repository=venkatesg559/cloud-service \
  --set image.tag=latest
CI/CD
On each push to main, GitHub Actions will:

Run Go tests

Build and push Docker image

Upgrade Helm release

Grafana Dashboard
Install Grafana in monitoring namespace

Enable the dashboards sidecar

Apply k8s/grafana-dashboard.yaml

View at http://localhost:3000 (admin/admin)

Contributing
Feel free to open issues or submit pull requests — happy to collaborate!

---

**How to add it**:

1. Create `README.md` in your repo root with the content above.
2. Commit and push:
   ```bash
   git add README.md
   git commit -m "docs: add project README"
   git push


