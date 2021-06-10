#!/usr/bin/env bash
# Teardown KinD cluster

if ! command -v kind &> /dev/null; then
  echo "KinD is not available"
  exit 1
fi

kind delete cluster
