---
slug: /zh-CN/credential-plugin/kubeconfig
title: get-kubeconfig（中文）
sidebar_position: 1
---

# get-kubeconfig

获取使用 ack-ram-tool 作为 [credential plugin](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins) 的 kubeconfig。

包含如下特性：

* 证书过期前将自动获取新的证书。
* 支持使用临时证书。
* 集成 ack-ram-authenticator。

## 使用示例

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

## 命令行参数

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

参数说明：

| 参数名称                   | 默认值                                            | 必需参数 | 说明                                                                                                                        |
|------------------------|------------------------------------------------|------|---------------------------------------------------------------------------------------------------------------------------|
| -c, --cluster-id       | 无                                              | 是    | 集群 ID                                                                                                                     |
| -m, --mode             | certificate                                    | 否    | kubeconfig 中的认证方法： `certificate` 表示证书认证，`ram-authenticator-token` 表示基于 ack-ram-authenticator 的 token 认证                   |
| --expiration           | 3h                                             | 否    | --mode 被设置为 `certificate` 时，通过这个参数设置证书过期时间。为 0 时表示不使用临时证书而是使用有效期更长的证书（过期时间由服务端自动确定）                                       |
| --private-address      | false                                          | 否    | 是否使用内网 api server 地址                                                                                                      |
| --api-version          | v1beta1                                        | 否    | 指定返回的数据中使用哪个版本的 apiVersion。v1beta1 表示 `client.authentication.k8s.io/v1beta1`，v1 表示 `client.authentication.k8s.io/v1beta1` |
| --credential-cache-dir | `~/.kube/cache/ack-ram-tool/credential-plugin` | 否    | 用于缓存证书的目录，只在 `--mode` 被设置为 `certificate` 时有效                                                                              |
