---
slug: /zh-CN/credential-plugin/get-token
title: get-token（中文）
sidebar_position: 3
---

# get-token

集成 ack-ram-authenticator，获取用于访问 api server 的 [ExecCredential](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins) token。

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
  "token": "k8s-ack-v1.aHR0cHM6Ly9zd***"
 }
}
```

## 命令行参数

```
Usage:
  ack-ram-tool credential-plugin get-token [flags]

Flags:
      --api-version string   v1 or v1beta1 (default "v1beta1")
  -c, --cluster-id string    The cluster id to use
  -h, --help                 help for get-token

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

| 参数名称             | 默认值     | 必需参数 | 说明                                                                                                                        |
|------------------|---------|------|---------------------------------------------------------------------------------------------------------------------------|
| -c, --cluster-id | 无       | 是    | 集群 ID                                                                                                                     |
| --api-version    | v1beta1 | 否    | 指定返回的数据中使用哪个版本的 apiVersion。v1beta1 表示 `client.authentication.k8s.io/v1beta1`，v1 表示 `client.authentication.k8s.io/v1beta1` |
