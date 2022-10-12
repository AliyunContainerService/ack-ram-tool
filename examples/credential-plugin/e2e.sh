#!/usr/bin/env bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"
CLUSTER_ID="$1"
KUBECONFIG_PATH="${SCRIPT_DIR}/kubeconfig"
CACHE_DIR="${HOME}/.kube/cache/ack-ram-tool/credential-plugin"

trap cleanup EXIT

function bar_tip() {
  echo -e "\n=== $1 ===\n"
}

function get_kubeconfig() {
  bar_tip "get kubeconfig"

  ack-ram-tool credential-plugin get-kubeconfig --cluster-id ${CLUSTER_ID} > ${KUBECONFIG_PATH}
}

function exec_auth() {
  bar_tip "exec auth plugin"

  kubectl --kubeconfig=${KUBECONFIG_PATH} get ns
}

function cleanup() {
  set +e
  bar_tip "cleanup"

  rm ${KUBECONFIG_PATH}
  rm ${CACHE_DIR}/*
  rm -r ${CACHE_DIR}

  set -e
}

function main() {
  if [[ "${CLUSTER_ID}none" == "none" ]]; then
    echo "clusterId is missing. Usage: bash e2e.sh CLUSTER_ID"
    exit 1
  fi
  if [[ "${SCRIPT_DIR}none" == "none" ]]; then
    echo "get script dir failed"
    exit 1
  fi

  get_kubeconfig
  exec_auth
}

main
