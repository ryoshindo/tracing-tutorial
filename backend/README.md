# backend

## setup

```shell
brew install minikube skaffold kustomize istioctl
```

## start server

```shell
minikube start --cpus=4 --memory=8gb
istioctl install -y --set profile=default
skaffold dev
minikube tunnel

curl localhost:80/ready # ready
```
