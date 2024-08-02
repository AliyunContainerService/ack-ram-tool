# rocketmq-4.x-java-sdk

Using [rocketmq-4.x-java-client](https://www.alibabacloud.com/help/doc-detail/445534.html) with RRSA Auth.

```
<dependency>
    <groupId>com.aliyun</groupId>
    <artifactId>credentials-java</artifactId>
    <version>0.3.5</version>
</dependency>
<dependency>
    <groupId>com.aliyun</groupId>
    <artifactId>tea</artifactId>
    <version>1.3.0</version>
</dependency>
```

https://github.com/aliyun/credentials-java


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


4. Associate an RAM Role to the service account and attach the policy to the role:

```
ack-ram-tool rrsa associate-role --cluster-id "${CLUSTER_ID}" \
    --namespace rrsa-demo-rocketmq-4x-java-sdk \
    --service-account demo-sa \
    --role-name test-rrsa-demo \
    --create-role-if-not-exist \
    --attach-system-policy AliyunMQPubOnlyAccess

ack-ram-tool rrsa associate-role --cluster-id "${CLUSTER_ID}" \
    --namespace rrsa-demo-rocketmq-4x-java-sdk \
    --service-account demo-sa \
    --role-name test-rrsa-demo \
    --create-role-if-not-exist \
    --attach-system-policy AliyunMQSubOnlyAccess
```

5. Deploy demo job:

```
export MQ_ENDPOINT="<MQ_ENDPOINT>"

ack-ram-tool credential-plugin get-kubeconfig --cluster-id "${CLUSTER_ID}" > kubeconfig
sed "s#__MQ_ENDPOINT__#${MQ_ENDPOINT}#g" deploy.yaml.tpl > deploy.yaml
kubectl --kubeconfig ./kubeconfig apply -f deploy.yaml
```

6. Get logs:

```
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-rocketmq-4x-java-sdk wait --for=condition=complete job/demo --timeout=240s
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-rocketmq-4x-java-sdk logs job/demo
```

Outputs:

```
test RocketMQ 4.x sdk use rrsa oidc token
====== Producer =======
SLF4J(W): No SLF4J providers were found.
SLF4J(W): Defaulting to no-operation (NOP) logger implementation
SLF4J(W): See https://www.slf4j.org/codes.html#noProviders for further details.
Producer Started.
[Producer][2024-08-02 10:26:04] SendResult msgId: 7F000001001E531******
====== Consumer =======
Consumer Started.
[Consumer][2024-08-02 10:26:04] Receive New Messages: [msgId: 7F000001001E531******, ]
```
