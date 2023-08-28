---
slug: get-kubeconfig
sidebar_position: 1
---

# get-kubeconfig

Obtain a kubeconfig file that uses ack-ram-tool as the [credential plugin](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins).

It has the following features:

* Automatically obtains a new certificate before the certificate expires.
* Supports using temporary certificate.
* Integrate ack-ram-authenticator.

## Usage

```shell
$ ack-ram-tool credential-plugin get-kubeconfig --cluster-id c5e***

kind: Config
apiVersion: v1
clusters:
    - name: kubernetes
      cluster:
        server: https://106.*.*.*:6443
        certificate-authority-data: LS0tL***
contexts:
    - name: 272***-c5e***
      context:
        cluster: kubernetes
        user: "272***"
current-context: 272***-c5e***
users:
    - name: "272***"
      user:
        exec:
            command: ack-ram-tool
            args:
                - credential-plugin
                - get-credential
                - --cluster-id
                - c5e***
                - --api-version
                - v1beta1
                - --expiration
                - 3h
                - --log-level
                - error
            apiVersion: client.authentication.k8s.io/v1beta1
            provideClusterInfo: false
            interactiveMode: Never
preferences: {}

$ ack-ram-tool credential-plugin get-kubeconfig --cluster-id c5e*** > kubeconfig
$ kubectl --kubeconfig kubeconfig get ns
NAME                         STATUS   AGE
default                      Active   6d3h
kube-node-lease              Active   6d3h
kube-public                  Active   6d3h
kube-system                  Active   6d3h
```

### --mode ram-authenticator-token

```
$ ack-ram-tool credential-plugin get-kubeconfig --mode ram-authenticator-token --cluster-id c5e***

kind: Config
apiVersion: v1
clusters:
    - name: kubernetes
      cluster:
        server: https://106.*.*.*:6443
        certificate-authority-data: LS0t***
contexts:
    - name: 272***-c5e***
      context:
        cluster: kubernetes
        user: "272***"
current-context: 272***-c5e***
users:
    - name: "272***"
      user:
        exec:
            command: ack-ram-tool
            args:
                - credential-plugin
                - get-token
                - --cluster-id
                - c5e***
                - --api-version
                - v1beta1
                - --log-level
                - error
            apiVersion: client.authentication.k8s.io/v1beta1
            provideClusterInfo: false
            interactiveMode: Never
preferences: {}

```

## Flags

```
Usage:
  ack-ram-tool credential-plugin get-kubeconfig [flags]

Flags:
      --api-version string            v1 or v1beta1 (default "v1beta1")
  -c, --cluster-id string             The cluster id to use
      --credential-cache-dir string   Directory to cache certificate (default "~/.kube/cache/ack-ram-tool/credential-plugin")
      --expiration duration           The certificate expiration (default 3h0m0s)
  -h, --help                          help for get-kubeconfig
  -m, --mode string                   credential mode: certificate or ram-authenticator-token (default "certificate")
      --private-address               Use private ip as api-server address

Global Flags:
  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively
      --ignore-aliyun-cli-credentials   don't try to parse credentials from config.json of aliyun cli
      --ignore-env-credentials          don't try to parse credentials from environment variables
      --log-level string                log level: info, debug, error (default "info")
      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)
      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli
      --region-id string                The region to use (default "cn-hangzhou")
```

Descriptions：

| Flag                   | Default                                        | Required | Description                                                                                                                                                                                                                                                       |
|------------------------|------------------------------------------------|----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| -c, --cluster-id       |                                                | Yes      | Cluster ID                                                                                                                                                                                                                                                        |
| -m, --mode             | certificate                                    |          | Authentication methods in kubeconfig: `certificate` indicates certificate authentication, and `ram-authenticator-token` indicates token authentication based on ack-ram-authenticator                                                                             |
| --expiration           | 3h                                             |          | When --mode is set to `certificate`, set the certificate expiration time through this parameter. When it is 0, it means not to use a temporary certificate but to use a longer valid certificate (the expiration time is automatically determined by the server). |
| --private-address      | false                                          |          | Whether to use the intranet API server address?                                                                                                                                                                                                                   |
| --api-version          | v1beta1                                        |          | Specify which version of apiVersion to use in the returned data. `v1beta1` represents `client.authentication.k8s.io/v1beta1`, and `v1` represents `client.authentication.k8s.io/v1`.                                                                              |
| --credential-cache-dir | `~/.kube/cache/ack-ram-tool/credential-plugin` |          | The directory used to cache the certificate is only valid when `--mode` is set to `certificate`                                                                                                                                                                   |
