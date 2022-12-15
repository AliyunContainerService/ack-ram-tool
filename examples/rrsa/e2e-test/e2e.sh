#!/usr/bin/env bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"
CLUSTER_ID="$1"
ROLE_NAME="test-rrsa-${CLUSTER_ID}"
POLICY_NAME="AliyunCSReadOnlyAccess"
KUBECONFIG_PATH="${SCRIPT_DIR}/kubeconfig"
NAMESPACE="test-rrsa"
SERVICE_ACCOUNT="sa-abc"
POD_NAME="demo"

trap cleanup EXIT

function bar_tip() {
  echo -e "\n=== $1 ===\n"
}

function enable_rrsa() {
  bar_tip "enable rrsa"
  ack-ram-tool rrsa -y -c "${CLUSTER_ID}" enable
  ack-ram-tool rrsa -y -c "${CLUSTER_ID}" status

  arn=$(ack-ram-tool rrsa -y -c "${CLUSTER_ID}" status|grep Arn  |awk '{print $4}')
  export OIDC_ARN=${arn}
}

function get_kubeconfig() {
  bar_tip "get and setup kubeconfig"
  aliyun cs DescribeClusterUserKubeconfig --ClusterId "${CLUSTER_ID}" --TemporaryDurationMinutes 15 \
    --endpoint cs.aliyuncs.com | jq '.config' -r > ${KUBECONFIG_PATH}
  export KUBECONFIG=${KUBECONFIG_PATH}
}

function create_resources() {
  bar_tip "create resources"
  set +e
  kubectl create ns ${NAMESPACE}
  kubectl create sa ${SERVICE_ACCOUNT} -n ${NAMESPACE}
  set -e
}

function associate_role() {
  bar_tip "associate role"
  ack-ram-tool rrsa -y -c "${CLUSTER_ID}" associate-role --create-role-if-not-exist \
                    -r ${ROLE_NAME} -n ${NAMESPACE} -s ${SERVICE_ACCOUNT}

  arn=$(aliyun ram GetRole --RoleName ${ROLE_NAME} |jq -r .Role.Arn)
  export ROLE_ARN=${arn}
}

function attach_policy_to_role() {
  bar_tip "attach policy to role"

  if aliyun ram ListPoliciesForRole --RoleName ${ROLE_NAME} | grep ${POLICY_NAME}; then
    return
  fi

  aliyun ram AttachPolicyToRole --PolicyType System --PolicyName ${POLICY_NAME} \
                                --RoleName ${ROLE_NAME}
}

function deploy_pod() {
  bar_tip "deploy pod"
  kubectl -n ${NAMESPACE} delete pod --all

  sed "s#__ALIBABA_CLOUD_ROLE_ARN__#${ROLE_ARN}#g" "${SCRIPT_DIR}/deploy.yaml" | \
    sed "s#__ALIBABA_CLOUD_OIDC_PROVIDER_ARN__#${OIDC_ARN}#g" | \
    kubectl -n ${NAMESPACE} apply -f -
}

function wait_pod_success() {
  bar_tip "wait pod success"
  kubectl -n ${NAMESPACE} wait --for=condition=Initialized pod/${POD_NAME} --timeout=240s
  sleep 30
  kubectl -n ${NAMESPACE} get pod ${POD_NAME} |grep Completed
  kubectl -n ${NAMESPACE} logs --tail 10 ${POD_NAME}
}

function test_setup_addon() {
  ack-ram-tool rrsa setup-addon --addon-name kritis-validation-hook \
    --cluster-id ${CLUSTER_ID} -y
}

function cleanup() {
  set +e
  bar_tip "cleanup"
  if ! echo "${SKIP}" |grep cleanup; then
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
  get_kubeconfig
  create_resources
  associate_role
  attach_policy_to_role
  deploy_pod
  wait_pod_success
  test_setup_addon
}

main
