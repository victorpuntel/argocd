Continuous Development using GO, ArgoCD, GitOps, Docker and Kubernetes.

Useful commands:
$ kind create cluster --name=argocd ## Create Kubernete's cluster
$ kubectl get nodes ## List clusters
$ kubectl get pods ## List pods
$ kubectl get svc ## List services

$ docker build -t victorpuntel/argocd:latest . ## Build docker image from Dockerfile on folder .
$ docker run --rm -p 8080:8080 victorpuntel/argocd:latest ## Run builded image
$ kubectl apply -f k8s/deployment.yaml ## Run deploy inside Kubernete's cluster
$ kubectl apply -f k8s/deployment.yaml ## Apply service in Kubernete's cluster
$ kubectl port-forward svc/goapp 8080:8080 ## Test app - forward port from service to computer
