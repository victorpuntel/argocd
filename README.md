
# Continuous Development with ArgoCD

A simple structure of CI/CD using GO, Docker, Kubernetes and GitOps.
### 🛠 Useful commands
Some useful commands to make it work

Kind
```bash
  $ kind create cluster --name=argocd ## Create Kubernete's cluster
```

Kubectl
```bash
  $ kubectl create namespace argocd 
  $ kubectl get nodes ## List clusters
  $ kubectl get pods ## List pods
  $ kubectl get svc ## List services
  $ kubectl apply -f k8s/deployment.yaml ## Run deploy inside Kubernete's cluster
  $ kubectl apply -f k8s/deployment.yaml ## Apply service in Kubernete's cluster
  $ kubectl port-forward svc/goapp 8080:8080 ## Test app - forward port from service to computer
```

Docker
```bash
  $ docker build -t victorpuntel/argocd:latest . ## Build docker image from Dockerfile on folder .
  $ docker run --rm -p 8080:8080 victorpuntel/argocd:latest ## Run builded image
```

ArgoCD
```bash
  $ kubectl create namespace argocd ## Needed to install ArgoCD
  $ kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml ## Needed to install ArgoCD
```