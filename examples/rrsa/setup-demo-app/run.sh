#!/usr/bin/env bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"
CLUSTER_ID="$1"
ROLE_NAME="demo-role-for-rrsa"
POLICY_NAME="AliyunCSReadOnlyAccess"
KUBECONFIG_PATH="${SCRIPT_DIR}/kubeconfig"
NAMESPACE="rrsa-demo"
SERVICE_ACCOUNT="demo-sa"

function bar_tip() {
  echo -e "\n=== $1 ===\n"
}

function enable_rrsa() {
  bar_tip "enable rrsa"
  ack-ram-tool rrsa -y -c "${CLUSTER_ID}" enable
  ack-ram-tool rrsa -y -c "${CLUSTER_ID}" status
}

function install_helper_addon() {
    bar_tip "install helper addon"
    ack-ram-tool rrsa -y -c "${CLUSTER_ID}" install-helper-addon
}

function get_kubeconfig() {
  bar_tip "get and setup kubeconfig"
  ack-ram-tool credential-plugin get-kubeconfig --cluster-id "${CLUSTER_ID}" > ${KUBECONFIG_PATH}
  export KUBECONFIG=${KUBECONFIG_PATH}
}

function associate_role_and_attach_policy() {
  bar_tip "associate role and attach policy"
  ack-ram-tool rrsa -y --cluster-id "${CLUSTER_ID}" associate-role \
      --role-name ${ROLE_NAME} --namespace ${NAMESPACE} --service-account ${SERVICE_ACCOUNT} \
      --create-role-if-not-exist --attach-system-policy ${POLICY_NAME}
}

function deploy_workload() {
  bar_tip "deploy workload"
  kubectl delete -f "${SCRIPT_DIR}/deploy.yaml" || true
  kubectl apply -f "${SCRIPT_DIR}/deploy.yaml"
}

function main() {
  if [[ "${CLUSTER_ID}none" == "none" ]]; then
    echo "clusterId is missing. Usage: bash run.sh CLUSTER_ID"
    exit 1
  fi
  if [[ "${SCRIPT_DIR}none" == "none" ]]; then
    echo "get script dir failed"
    exit 1
  fi

  enable_rrsa
  install_helper_addon
  get_kubeconfig
  associate_role_and_attach_policy
  deploy_workload
}

main
