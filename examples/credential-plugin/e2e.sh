#!/usr/bin/env bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"
CLUSTER_ID="$1"
KUBECONFIG_PATH="${SCRIPT_DIR}/kubeconfig"
CACHE_DIR="${HOME}/.kube/cache/ack-ram-tool/credential-plugin"
MODE="$2"

trap cleanup EXIT

function bar_tip() {
  echo -e "\n=== $1 ===\n"
}

function get_kubeconfig() {
  bar_tip "get kubeconfig"

  ack-ram-tool credential-plugin get-kubeconfig -m ${MODE} --cluster-id ${CLUSTER_ID} > ${KUBECONFIG_PATH}

  if echo ${MODE} |grep token; then
    Arn=$(aliyun sts GetCallerIdentity | jq .Arn -r)
    UserId=$(aliyun sts GetCallerIdentity | jq .UserId -r)
    ack-ram-tool credential-plugin get-kubeconfig --cluster-id ${CLUSTER_ID} > ${KUBECONFIG_PATH}.crt.yaml
    cat <<EOF | kubectl --kubeconfig=${KUBECONFIG_PATH}.crt.yaml apply -f -
apiVersion: ramauthenticator.k8s.alibabacloud/v1alpha1
kind: RAMIdentityMapping
metadata:
  name: "${UserId}"
spec:
  arn: ${Arn}
  username: "${UserId}"
EOF
  fi
}

function exec_auth() {
  bar_tip "exec auth plugin"

  kubectl --kubeconfig=${KUBECONFIG_PATH} get ns
  kubectl --kubeconfig=${KUBECONFIG_PATH} auth whoami
}

function cleanup() {
  set +e
  bar_tip "cleanup"

  rm ${KUBECONFIG_PATH}
  rm ${KUBECONFIG_PATH}.crt.yaml
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
