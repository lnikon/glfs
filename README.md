# Graph Library Provider Service

## Use docker for development and deployment

To use the graph library frontend-server(glfs) you cun lunch docker images in the following way:

```
% docker run --rm -it -d -v $(pwd):/app -p 8090:8090 --name glfs-docker gl:latest
```

To lunch deploy the glfs using Kubernetes engine the deployment is provided.
Note, that if your testing the glfs locally using minikube cluster, then minikube tunnel should be lunched
as minikube does not support load balancing to forward the requests to the pods.

## ToDo

- Add arguments parsing
- Containerize development
- Support linting
- Support pipeline
- Deploy to GKE
- Add endpoint testing
- Integrate graph serialization
