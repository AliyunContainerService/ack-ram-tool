# go-sdk

Using [Alibaba Could Go SDK](https://github.com/aliyun/alibabacloud-go-sdk) with RRSA Auth.

```
go get github.com/aliyun/credentials-go@v1.2.7
```

https://github.com/aliyun/credentials-go


## Demo

1. Enable RRSA:

```
export CLUSTER_ID=<cluster_id>
ack-ram-tool rrsa enable --cluster-id "${CLUSTER_ID}"
```

2. Install ack-pod-identity-webhook:

```
ack-ram-tool rrsa install-helper-addon --cluster-id "${CLUSTER_ID}"
```

3. Create an RAM Policy:

```
aliyun ram CreatePolicy --PolicyName cs-describe-clusters --PolicyDocument '{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "cs:DescribeClusters",
        "cs:GetClusters"
      ],
      "Resource": [
        "*"
      ],
      "Condition": {}
    }
  ]
}'
```

4. Associate an RAM Role to the service account and attach the policy to the role:

```
ack-ram-tool rrsa associate-role --cluster-id "${CLUSTER_ID}" \
    --namespace rrsa-demo-golang-sdk \
    --service-account demo-sa \
    --role-name test-rrsa-demo \
    --create-role-if-not-exist \
    --attach-custom-policy cs-describe-clusters
```

5. Deploy demo job:

```
ack-ram-tool credential-plugin get-kubeconfig --cluster-id "${CLUSTER_ID}" > kubeconfig
kubectl --kubeconfig ./kubeconfig apply -f deploy.yaml
```

6. Get logs:

```
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-golang-sdk wait --for=condition=complete job/demo --timeout=240s
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-golang-sdk logs job/demo
```

Outputs:

```
test open api sdk use rrsa oidc token

cluster id: c4db8***, cluster name: foo***
cluster id: cc20c***, cluster name: bar***

```
