#1 /usr/bin/env sh

set -e

# wait for shopper_service
kubectl rollout status deployment/podinfo --timeout=3m

# test shopper_service
helm test podinfo
