---
slug: /zh-CN/credential-plugin/get-credential
title: get-credential（中文）
sidebar_position: 2
---

# get-credential

获取用于访问 api server 的 ExecCredential 证书数据。

包含如下特性：

* 证书过期前将自动获取新的证书
* 支持使用临时证书


## 使用示例

```shell
$ ack-ram-tool credential-plugin get-token --cluster-id <clusterId>

{
 "kind": "ExecCredential",
 "apiVersion": "client.authentication.k8s.io/v1beta1",
 "spec": {
  "interactive": false
 },
 "status": {
  "expirationTimestamp": "2023-04-20T09:29:06Z",
  "clientCertificateData": "-----BEGIN CERTIFICATE-----\nMIID***\n-----END CERTIFICATE-----\n",
  "clientKeyData": "-----BEGIN RSA PRIVATE KEY-----\nMIIE***\n-----END RSA PRIVATE KEY-----\n"
 }
}
```

## 命令行参数

```
Usage:
  ack-ram-tool credential-plugin get-credential [flags]

Flags:
      --api-version string            v1 or v1beta1 (default "v1beta1")
  -c, --cluster-id string             The cluster id to use
      --credential-cache-dir string   Directory to cache credential (default "~/.kube/cache/ack-ram-tool/credential-plugin")
      --expiration duration           The credential expiration (default 3h0m0s)
  -h, --help                          help for get-credential

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

| 参数名称             | 默认值                                            | 必需参数 | 说明                                                                                                                        |
|------------------|------------------------------------------------|------|---------------------------------------------------------------------------------------------------------------------------|
| -c, --cluster-id | 无                                              | 是    | 集群 ID                                                                                                                     |
| --api-version    | v1beta1                                        | 否    | 指定返回的数据中使用哪个版本的 apiVersion。v1beta1 表示 `client.authentication.k8s.io/v1beta1`，v1 表示 `client.authentication.k8s.io/v1beta1` |
| --expiration    | 3h0m0s                                         | 否    | 指定证书过期时间。为 0 时表示不使用临时证书而是使用有效期更长的证书（过期时间由服务端自动确定）                                                                         |
| --credential-cache-dir    | `~/.kube/cache/ack-ram-tool/credential-plugin` | 否    | 用于缓存证书的目录                                                                                                                 |
