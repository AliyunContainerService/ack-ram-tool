# aliyunlog cli demo

## Usage

1. Enable RRSA:

```
export CLUSTER_ID=<cluster_id>
ack-ram-tool rrsa enable --cluster-id "${CLUSTER_ID}"
```

2. Install ack-pod-identity-webhook:

```
ack-ram-tool rrsa install-helper-addon --cluster-id "${CLUSTER_ID}"
```

3. Create a RAM Role and attach a system policy to the role:

```
ack-ram-tool rrsa associate-role --cluster-id "${CLUSTER_ID}" \
    --namespace rrsa-demo-aliyunlog-cli \
    --service-account demo-sa \
    --role-name test-rrsa-demo \
    --create-role-if-not-exist \
    --attach-system-policy AliyunLogReadOnlyAccess
```

4. Deploy demo job:

```
ack-ram-tool credential-plugin get-kubeconfig --cluster-id "${CLUSTER_ID}" > kubeconfig
kubectl --kubeconfig ./kubeconfig apply -f deploy.yaml
```

5. Get logs:

```
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-aliyunlog-cli wait --for=condition=complete job/demo --timeout=240s
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-aliyunlog-cli logs job/demo
```

Outputs:

```
{"count": 1, "projects": [
{"createTime": "1676282996", "description": "k8s log project, ***",
 "lastModifyTime": "1676282996", "owner": "", "projectName": "k8s-log-c0edc***", 
 "region": "cn-hangzhou", "resourceGroupId": "rg-***", "status": "Normal"}],
  "total": 24}
```
