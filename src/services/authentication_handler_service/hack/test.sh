#1 /usr/bin/env sh

set -e

# wait for authentication_handler_service
kubectl rollout status deployment/auth-hdlr-svc-deployment-authentication-handler-service --timeout=3m

# test authentication_handler_service
helm test auth-hdlr-svc-deployment
