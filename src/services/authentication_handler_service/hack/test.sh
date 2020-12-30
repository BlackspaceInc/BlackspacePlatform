#1 /usr/bin/env sh

set -e

# wait for authentication_handler_service
kubectl rollout status deployment/authsvcployment-authentication-handler-service --timeout=3m

# test authentication_handler_service
helm test authsvcployment-authentication-handler-service
