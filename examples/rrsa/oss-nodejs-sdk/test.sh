#!/usr/bin/env bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"
CLUSTER_ID="$1"
KUBECONFIG_PATH="${SCRIPT_DIR}/kubeconfig"
NAMESPACE="rrsa-demo-oss-nodejs-sdk"
ROLE_NAME="test-rrsa-demo"
POLICY_NAME="test-oss-list-buckets"

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

  aliyun ram DeletePolicy --PolicyName ${POLICY_NAME} || true
  aliyun ram CreatePolicy --PolicyName ${POLICY_NAME} --PolicyDocument '{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "oss:ListBuckets"
      ],
      "Resource": [
        "*"
      ],
      "Condition": {}
    }
  ]
}' || true

  ack-ram-tool rrsa associate-role --cluster-id "${CLUSTER_ID}" \
    --namespace "${NAMESPACE}" \
    --service-account demo-sa \
    --role-name ${ROLE_NAME} \
    --create-role-if-not-exist \
    --attach-custom-policy ${POLICY_NAME}
}

function deploy_demo() {
  bar_tip "deploy demo"

  ack-ram-tool credential-plugin get-kubeconfig --cluster-id "${CLUSTER_ID}" > ${KUBECONFIG_PATH}
  kubectl --kubeconfig ${KUBECONFIG_PATH} delete -f "${SCRIPT_DIR}/deploy.yaml" || true
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
  aliyun ram DetachPolicyFromRole --RoleName ${ROLE_NAME} --PolicyName ${POLICY_NAME} --PolicyType Custom || true

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
