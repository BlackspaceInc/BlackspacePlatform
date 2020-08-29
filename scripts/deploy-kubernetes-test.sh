#!/usr/bin/env bash
set -euo pipefail

# deploy-kubernetes-test.sh
#
# SUMMARY
#
#   Deploys Blackspace into Kubernetes for testing purposes.
#   Uses the same installation method our users would use.
#
#   This script implements cli interface required by the kubernetes E2E
#   tests.
#
# USAGE
#
#   Deploy:
#
#   $ CONTAINER_IMAGE=timberio/Blackspace:alpine-latest scripts/deploy-kubernetes-test.sh up Blackspace-test-qwerty
#
#   Teardown:
#
#   $ scripts/deploy-kubernetes-test.sh down Blackspace-test-qwerty
#

# Command to perform.
COMMAND="${1:?"Specify the command (up/down) as the first argument"}"

# A Kubernetes namespace to deploy to.
NAMESPACE="${2:?"Specify the namespace as the second argument"}"

# Allow overriding kubectl with somethingl like `minikube kubectl --`.
Blackspace_TEST_KUBECTL="${Blackspace_TEST_KUBECTL:-"kubectl"}"

# Allow optionally installing custom resource configs.
CUSTOM_RESOURCE_CONIFGS_FILE="${CUSTOM_RESOURCE_CONIFGS_FILE:-""}"


# TODO: replace with `helm template | kubectl apply -f -` when Helm Chart is
# available.

templated-config-global() {
  sed "s|^    namespace: Blackspace|    namespace: $NAMESPACE|" < "distribution/kubernetes/Blackspace-global.yaml" \
    | sed "s|^  name: Blackspace|  name: $NAMESPACE|"
}

up() {
  # A Blackspace container image to use.
  CONTAINER_IMAGE="${CONTAINER_IMAGE:?"You must assing CONTAINER_IMAGE variable with the Blackspace container image name"}"

  templated-config-global | $Blackspace_TEST_KUBECTL create -f -

  $Blackspace_TEST_KUBECTL create namespace "$NAMESPACE"

  if [[ -n "$CUSTOM_RESOURCE_CONIFGS_FILE" ]]; then
    $Blackspace_TEST_KUBECTL create --namespace "$NAMESPACE" -f "$CUSTOM_RESOURCE_CONIFGS_FILE"
  fi

  sed 's|image: timberio/Blackspace:[^$]*$'"|image: $CONTAINER_IMAGE|" < "distribution/kubernetes/Blackspace-namespaced.yaml" \
    | $Blackspace_TEST_KUBECTL create --namespace "$NAMESPACE" -f -
}

down() {
  # A workaround for `kubectl` from a `snap` package.
  cat < "distribution/kubernetes/Blackspace-namespaced.yaml" | $Blackspace_TEST_KUBECTL delete --namespace "$NAMESPACE" -f -

  if [[ -n "$CUSTOM_RESOURCE_CONIFGS_FILE" ]]; then
    $Blackspace_TEST_KUBECTL delete --namespace "$NAMESPACE" -f "$CUSTOM_RESOURCE_CONIFGS_FILE"
  fi

  templated-config-global | $Blackspace_TEST_KUBECTL delete -f -

  $Blackspace_TEST_KUBECTL delete namespace "$NAMESPACE"
}

case "$COMMAND" in
  up|down)
    "$COMMAND" "$@"
    ;;
  *)
    echo "Invalid command: $COMMAND" >&2
    exit 1
esac
