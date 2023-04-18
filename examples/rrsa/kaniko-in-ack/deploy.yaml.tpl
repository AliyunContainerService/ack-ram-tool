---
apiVersion: v1
kind: Namespace
metadata:
  name: rrsa-demo-kaniko
  labels:
    pod-identity.alibabacloud.com/injection: 'on'

---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: demo-sa
  namespace: rrsa-demo-kaniko
  annotations:
    pod-identity.alibabacloud.com/role-name: test-rrsa-demo

---
apiVersion: v1
kind: ConfigMap
metadata:
  namespace: rrsa-demo-kaniko
  name: docker-config
data:
  config.json: |
      {
        "credHelpers": {
          "REGISTRY_DOMAIN": "acr-helper"
        }
      }

---
apiVersion: batch/v1
kind: Job
metadata:
  name: demo
  namespace: rrsa-demo-kaniko
spec:
  template:
    spec:
      serviceAccountName: demo-sa
      restartPolicy: Never
      containers:
        - image: registry.cn-hangzhou.aliyuncs.com/acs/ack-ram-tool:0.13.0-dev-rrsa-example-kaniko
          args:
            - "--dockerfile=/app/Dockerfile"
            - "--context=dir:///app"
            - "--destination=DEST_IMAGE"
          imagePullPolicy: "Always"
          name: test
          volumeMounts:
            - name: docker-config
              mountPath: /kaniko/.docker/
      volumes:
        - name: docker-config
          configMap:
            name: docker-config
