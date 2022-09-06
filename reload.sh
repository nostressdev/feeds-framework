!/bin/bash
kubectl delete -f deployment/feeds-framework.yaml
kubectl apply -f deployment/feeds-framework.yaml
