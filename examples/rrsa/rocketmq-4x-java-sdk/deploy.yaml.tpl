---
apiVersion: v1
kind: Namespace
metadata:
  name: rrsa-demo-rocketmq-4x-java-sdk
  labels:
    pod-identity.alibabacloud.com/injection: 'on'

---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: demo-sa
  namespace: rrsa-demo-rocketmq-4x-java-sdk
  annotations:
    pod-identity.alibabacloud.com/role-name: test-rrsa-demo

---
apiVersion: batch/v1
kind: Job
metadata:
  name: demo
  namespace: rrsa-demo-rocketmq-4x-java-sdk
spec:
  template:
    spec:
      serviceAccountName: demo-sa
      restartPolicy: Never
      containers:
        - image: registry.cn-hangzhou.aliyuncs.com/acs/ack-ram-tool:0.14.0-rrsa-example-rocketmq-4x-java-sdk
          imagePullPolicy: "Always"
          name: test
          env:
            - name: MQ_ENDPOINT
              value: __MQ_ENDPOINT__
