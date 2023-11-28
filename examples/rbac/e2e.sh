#!/usr/bin/env bash
set -e

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"
CLUSTER_ID="$1"

trap cleanup EXIT

function bar_tip() {
  echo -e "\n=== $1 ===\n"
}

function scan_permissions_deleted_users() {
  bar_tip "scan-user-permissions"

  ack-ram-tool rbac scan-user-permissions --cluster-id ${CLUSTER_ID}
}

function scan_permissions_all_users() {
  bar_tip "scan-user-permissions all users"

  ack-ram-tool rbac scan-user-permissions --cluster-id ${CLUSTER_ID} --all-users
}

function scan_permissions_all_clusters() {
  bar_tip "scan-user-permissions all clusters"

  ack-ram-tool rbac scan-user-permissions --cluster-id all
}

function scan_permissions_all_clusters_all_users() {
  bar_tip "scan-user-permissions all clusters and all users"

  ack-ram-tool rbac scan-user-permissions --cluster-id all --all-users
}

function cleanup() {
  set +e
  bar_tip "cleanup"

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

  scan_permissions_deleted_users
  scan_permissions_all_users
  scan_permissions_all_clusters
  scan_permissions_all_clusters_all_users
}

main
