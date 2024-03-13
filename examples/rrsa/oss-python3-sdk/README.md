# oss-python3-sdk

Using [OSS Python 3 SDK](https://github.com/aliyun/aliyun-oss-python-sdk) with RRSA Auth.

```
pip install 'alibabacloud_credentials>=0.3.1'
```

https://github.com/aliyun/credentials-python


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
aliyun ram CreatePolicy --PolicyName oss-list-buckets --PolicyDocument '{
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
}'
```

4. Associate an RAM Role to the service account and attach the policy to the role:

```
ack-ram-tool rrsa associate-role --cluster-id "${CLUSTER_ID}" \
    --namespace rrsa-demo-oss-python3-sdk \
    --service-account demo-sa \
    --role-name test-rrsa-demo \
    --create-role-if-not-exist \
    --attach-custom-policy oss-list-buckets
```

5. Deploy demo job:

```
ack-ram-tool credential-plugin get-kubeconfig --cluster-id "${CLUSTER_ID}" > kubeconfig
kubectl --kubeconfig ./kubeconfig apply -f deploy.yaml
```

6. Get logs:

```
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-oss-python3-sdk wait --for=condition=complete job/demo --timeout=240s
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-oss-python3-sdk logs job/demo
```

Outputs:

```
2023/05/19 10:58:55 test oss sdk using rrsa oidc token
call oss.listBuckets via oidc token success:
- test-***
- cri-***

```
