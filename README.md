# Deploy Golang App and MYSQL Using Docker and Kubernetes

## Instalation

Pull golang app from [docker-hub](https://hub.docker.com/repository/docker/danisbagus/app-go)

```
docker push danisbagus/app-go:latest
```

Login to dockerhub account in terminal

```
docker login
```

Run Minikube

```
minikube start
```

Create Persistant Volume Claim, deployment, and service for Mysql

```
kubectl apply -f /k8s/mysql/.
```

Create Persistant Configmap, deployment, and service for Golang App

```
kubectl apply -f /k8s/.
```

## Verify

Verify Persistant Volume Claim

```
kubectl get pvc
```

Verify Deployment

```
kubectl get deployment
```

Verify Pod

```
kubectl get deployment
```

Verify Service

```
kubectl get service
```

Verify configmap

```
kubectl get configmap
```

## Access the service

```
minikube service <service name>
```

## Get an interactive shell on a pod + container

```
kubectl exec -it <pod name> -- /bin/bash
```

## Restart the deployment

```
kubectl scale deployment <deployment name> --replicas=0 -n default
```

```
kubectl scale deployment <deployment name> --replicas=1 -n default
```
