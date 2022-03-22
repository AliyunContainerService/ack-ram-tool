#!/usr/bin/env bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"
CLUSTER_ID="$1"
ROLE_NAME="test-rrsa-${CLUSTER_ID}"
KUBECONFIG_PATH="${SCRIPT_DIR}/kubeconfig"
NAMESPACE="test-rrsa"
SERVICE_ACCOUNT="user1"

trap cleanup EXIT

function bar_tip() {
  echo -e "\n=== $1 ===\n"
}

function enable_rrsa() {
  bar_tip "enable rrsa"
  ack-ram-tool rrsa -y -c "${CLUSTER_ID}" enable
  ack-ram-tool rrsa -y -c "${CLUSTER_ID}" status
}

function get_metadata() {
  bar_tip "get metadata"
  REGION=$(aliyun cs DescribeClusterDetail --ClusterId ${CLUSTER_ID} --endpoint cs.aliyuncs.com |jq '.region_id' -r)
  echo ${REGION}
  export REGION=${REGION}

  aliuid=$(aliyun sts GetCallerIdentity |jq -r .AccountId)
  export ALIUID=${aliuid}
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
  aliyun ram CreateRole --RoleName ${ROLE_NAME} --AssumeRolePolicyDocument \
    '{"Version": "1", "Statement": [{"Action": "sts:AssumeRole", "Effect": "Allow", "Principal": {"Service": ["cs.aliyuncs.com"]}}]}'
  set -e
}

function associate_role() {
  bar_tip "associate role"
  ack-ram-tool rrsa -y -c "${CLUSTER_ID}" associate-role -r ${ROLE_NAME} -n ${NAMESPACE} -s ${SERVICE_ACCOUNT}
}

function deploy_pods() {
  bar_tip "deploy pods"
  set +e
  kubectl -n ${NAMESPACE} delete pod --all
  set -e
  sed "s/REGION/${REGION}/g" "${SCRIPT_DIR}/deploy.yaml" | kubectl -n ${NAMESPACE} apply -f -
}

function assume_role() {
  bar_tip "assume role via oidc token"
  for name in $(echo run-as-root run-as-non-root); do
    kubectl -n ${NAMESPACE} wait --for=condition=Ready pod/${name} --timeout=240s
    TOKEN=$(kubectl -n ${NAMESPACE} exec -it ${name} -- cat /var/run/secrets/tokens/oidc-token)

    echo "assume-role via token from pod ${name}"
    echo ${TOKEN} | ack-ram-tool rrsa assume-role --region-id ${REGION} -r acs:ram::${ALIUID}:role/${ROLE_NAME} \
                      -p acs:ram::${ALIUID}:oidc-provider/ack-rrsa-${CLUSTER_ID} -t -
    echo ${REGION}
    echo $name
  done
}

function cleanup() {
  set +e
  bar_tip "cleanup"
  aliyun ram DeleteRole --RoleName ${ROLE_NAME}
  rm ${KUBECONFIG_PATH}
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

  get_metadata
  enable_rrsa
  get_kubeconfig
  create_resources
  associate_role
  deploy_pods
  assume_role
}

main