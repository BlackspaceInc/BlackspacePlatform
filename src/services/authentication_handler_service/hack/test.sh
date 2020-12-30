#1 /usr/bin/env sh

set -e

# wait for authentication_handler_service
kubectl rollout status deployment/authentication-handler-service --timeout=3m

# test authentication_handler_service
helm test authentication-handler-service
