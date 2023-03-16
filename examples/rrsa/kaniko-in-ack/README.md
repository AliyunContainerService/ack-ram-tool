# Running kaniko in ACK

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
    --namespace rrsa-demo-kaniko \
    --service-account demo-sa \
    --role-name test-rrsa-demo \
    --create-role-if-not-exist \
    --attach-system-policy AliyunContainerRegistryFullAccess
```

4. Deploy demo job:

```
export DEST_IMAGE="<image>"
ack-ram-tool credential-plugin get-kubeconfig --cluster-id "${CLUSTER_ID}" > kubeconfig
sed "s#DEST_IMAGE#${DEST_IMAGE}#g" deploy.yaml.tpl | \
    sed "s#REGISTRY_DOMAIN#`echo ${DEST_IMAGE}| cut -d '/' -f 1`#g"> deploy.yaml
kubectl --kubeconfig ./kubeconfig apply -f deploy.yaml
```

5. Get logs:

```
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-kaniko wait --for=condition=complete job/demo --timeout=240s
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-kaniko logs job/demo
```

Outputs:

```
INFO[0000] Retrieving image manifest alpine:3.17        
INFO[0000] Retrieving image alpine:3.17 from registry index.docker.io 
INFO[0003] Built cross stage deps: map[]                
INFO[0003] Retrieving image manifest alpine:3.17        
INFO[0003] Returning cached image manifest              
INFO[0003] Executing 0 build triggers                   
INFO[0003] Building stage 'alpine:3.17' [idx: '0', base-idx: '-1'] 
INFO[0003] Skipping unpacking as no commands require it. 
INFO[0003] ENTRYPOINT ["/bin/sh", "-c", "echo this image was build via Kaniko in ACK"] 
INFO[0003] Pushing image to registry.***.aliyuncs.com/***/***:v1 
INFO[0006] Pushed registry.***.aliyuncs.com/***/***@sha256:195104*** 
```
