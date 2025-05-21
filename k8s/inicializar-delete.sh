#!/bin/bash

set -e  # Detiene el script si hay un error

echo "📁 Desplegando microservicio p-go-delete..."

# Namespace
kubectl apply -f k8s/delete/namespace-delete.yaml

# Secret
kubectl apply -f k8s/delete/secrets-delete.yaml

# Deployment
kubectl apply -f k8s/delete/deployment-delete.yaml

# Esperar a que el deployment esté disponible
echo "⏳ Esperando a que p-go-delete esté listo..."
kubectl wait --namespace=p-go-delete \
  --for=condition=available deployment/delete-deployment \
  --timeout=90s

# Service
kubectl apply -f k8s/delete/service-delete.yaml

# Ingress
kubectl apply -f k8s/delete/ingress.yaml

echo "✅ p-go-delete desplegado correctamente."

echo -e "\n🔍 Estado actual:"
kubectl get all -n p-go-delete
kubectl get ingress -n p-go-delete
