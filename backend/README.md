# backend

## setup

```shell
brew install minikube skaffold kustomize istioctl
```

## start server

```shell
minikube start
skaffold dev
istioctl install --context minikube
```
