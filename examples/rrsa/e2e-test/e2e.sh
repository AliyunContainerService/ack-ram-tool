#!/usr/bin/env bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"
CLUSTER_ID="$1"
ROLE_NAME="test-rrsa-demo"
POLICY_NAME="AliyunCSReadOnlyAccess"
KUBECONFIG_PATH="${SCRIPT_DIR}/kubeconfig"
NAMESPACE="rrsa-demo"
SERVICE_ACCOUNT="demo-sa"
JOB_NAME="demo"

trap cleanup EXIT

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

function wait_pod_success() {
  bar_tip "wait pod success"
  kubectl -n ${NAMESPACE} wait --for=condition=complete job/${JOB_NAME} --timeout=240s
  kubectl -n ${NAMESPACE} logs --tail 10 job/${JOB_NAME}
}

function test_setup_addon() {
  ack-ram-tool rrsa setup-addon --addon-name kritis-validation-hook \
    --cluster-id ${CLUSTER_ID} -y
}

function cleanup() {
  set +e
  bar_tip "cleanup"
  if ! echo "${SKIP}" |grep cleanup; then
    aliyun ram DetachPolicyFromRole --RoleName ${ROLE_NAME} \
              --PolicyName ${POLICY_NAME} --PolicyType "System"
    aliyun ram DeleteRole --RoleName ${ROLE_NAME}

    aliyun ram DetachPolicyFromRole --RoleName "kritis-validation-hook-${CLUSTER_ID}" \
          --PolicyName "ack-kritis-validation-hook" --PolicyType "Custom"
    aliyun ram DeleteRole --RoleName "kritis-validation-hook-${CLUSTER_ID}"
    rm ${KUBECONFIG_PATH}
  fi
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

  enable_rrsa
  install_helper_addon
  get_kubeconfig
  associate_role_and_attach_policy
  deploy_workload
  wait_pod_success
  test_setup_addon
}

main
