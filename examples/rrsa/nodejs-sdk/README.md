# nodejs-sdk

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

3. Create a RAM Role:

```
ack-ram-tool rrsa associate-role --cluster-id "${CLUSTER_ID}" \
    --namespace rrsa-demo-nodejs-sdk \
    --service-account demo-sa \
    --role-name test-rrsa-demo \
    --create-role-if-not-exist \
    --attach-system-policy AliyunCSReadOnlyAccess
```

4. Deploy demo job:

```
ack-ram-tool credential-plugin get-kubeconfig --cluster-id "${CLUSTER_ID}" > kubeconfig
kubectl --kubeconfig ./kubeconfig apply -f deploy.yaml
```

5. Get logs:

```
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-nodejs-sdk wait --for=condition=complete job/demo --timeout=240s
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-nodejs-sdk logs job/demo
```

Outputs:

```
> nodejs-sdk@1.0.0 demo
> node_modules/.bin/ts-node src/index.ts

cluster id: c4db8***, cluster name: foo***
cluster id: cc20c***, cluster name: bar***

```
