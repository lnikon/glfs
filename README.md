# Graph Library Provider Service

## Use docker for development and deployment

```
% docker run --rm -it -d -v $(pwd):/app -p 8090:8090 --name glfs-docker gl:latest 
```

## ToDo

- Add arguments parsing
- Containerize development
- Support linting
- Support pipeline
- Deploy to GKE
- Add endpoint testing
- Integrate graph serialization
