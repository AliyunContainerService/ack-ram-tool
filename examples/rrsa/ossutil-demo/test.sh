#!/usr/bin/env bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"
CLUSTER_ID="$1"
KUBECONFIG_PATH="${SCRIPT_DIR}/kubeconfig"
NAMESPACE="rrsa-demo-ossutil"

trap cleanup EXIT

function bar_tip() {
  echo -e "\n=== $1 ===\n"
}

function enable_rrsa() {
  bar_tip "enable RRSA"

  ack-ram-tool rrsa enable --cluster-id "${CLUSTER_ID}"
}

function install_helper() {
  bar_tip "install ack-pod-identity-webhook"

  ack-ram-tool rrsa install-helper-addon --cluster-id "${CLUSTER_ID}"
}

function setup_role() {
  bar_tip "setup ram role"

  ack-ram-tool rrsa associate-role --cluster-id "${CLUSTER_ID}" \
    --namespace "${NAMESPACE}" \
    --service-account demo-sa \
    --role-name test-rrsa-demo \
    --create-role-if-not-exist \
    --attach-system-policy AliyunCSReadOnlyAccess
}

function deploy_demo() {
  bar_tip "deploy demo"

  ack-ram-tool credential-plugin get-kubeconfig --cluster-id "${CLUSTER_ID}" > ${KUBECONFIG_PATH}
  kubectl --kubeconfig ${KUBECONFIG_PATH} apply -f "${SCRIPT_DIR}/deploy.yaml"
}

function get_logs() {
  bar_tip "wait demo and get logs"

  kubectl --kubeconfig ${KUBECONFIG_PATH} -n "${NAMESPACE}" wait --for=condition=complete job/demo --timeout=240s
  kubectl --kubeconfig ${KUBECONFIG_PATH} -n "${NAMESPACE}" logs job/demo
}

function cleanup() {
  set +e
  bar_tip "cleanup"

  rm ${KUBECONFIG_PATH}

  set -e
}

function main() {
  if [[ "${CLUSTER_ID}none" == "none" ]]; then
    echo "clusterId is missing. Usage: bash test.sh CLUSTER_ID"
    exit 1
  fi
  if [[ "${SCRIPT_DIR}none" == "none" ]]; then
    echo "get script dir failed"
    exit 1
  fi

  enable_rrsa
  install_helper
  setup_role
  sleep 60
  deploy_demo
  get_logs
}

main
