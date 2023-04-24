---
slug: /zh-CN/rrsa/associate-role
title: associate-role（中文）
sidebar_position: 2
---

# associate-role

配置 RAM 角色，允许使用表示特定 service account 身份的 oidc token 扮演该 RAM 角色。

## 使用示例

```shell
$ ack-ram-tool rrsa associate-role --cluster-id <clusterId> \
  --namespace <namespce> --service-account <serviceAccountName> \
  --role-name <roleName>

? Are you sure you want to associate RAM Role "<roleName>" to service account "<serviceAccountName>" (namespace: "<namespce>")? Yes
2023-04-20T14:30:02+08:00 INFO will change the AssumeRole Policy of RAM Role "<roleName>" with blow content:
{
  "Statement": [
   {
    "Action": "sts:AssumeRole",
    "Condition": {
     "StringEquals": {
      "oidc:aud": "sts.aliyuncs.com",
      "oidc:iss": "https://oidc-ack-***.aliyuncs.com/c132c***",
      "oidc:sub": "system:serviceaccount:<namespce>:<serviceAccountName>"
     }
    },
    "Effect": "Allow",
    "Principal": {
     "Federated": [
      "acs:ram::113***:oidc-provider/ack-rrsa-c132c***"
     ]
    }
   }
  ],
  "Version": "1"
 }

? Are you sure you want to associate RAM Role "test" to service account "sa" (namespace: "test")? Yes
2023-04-20T14:30:04+08:00 INFO Associate RAM Role "test" to service account "sa" (namespace: "test") successfully
```

## 命令行参数

```
Usage:
  ack-ram-tool rrsa associate-role [flags]

Flags:
      --attach-custom-policy string   Attach this custom policy to the RAM Role
      --attach-system-policy string   Attach this system policy to the RAM Role
  -c, --cluster-id string             The cluster id to use
      --create-role-if-not-exist      Create the RAM Role if it does not exist
  -h, --help                          help for associate-role
  -n, --namespace string              The Kubernetes namespace to use
  -r, --role-name string              The RAM Role name to use
  -s, --service-account string        The Kubernetes service account to use

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

| 参数名称                       | 默认值 | 必需参数             | 说明                                       |
|----------------------------|-----|------------------|------------------------------------------|
| -c, --cluster-id           | 无   | 是                | 集群 ID                                    |
| -n, --namespace            | 无   | 是                | 命名空间，可以使用 `*` 表示所有命名空间：`--namespace '*'` |
| -s, --service-account      | 无   | 是                | service account                          |
| -r, --role-name            | 无   | 是                | RAM 角色                                   |
| --create-role-if-not-exist | 无   | 否                | 如果该 RAM 角色不存在，那么自动创建一个同名的 RAM 角色         |
| --attach-system-policy     | 无   | 为该角色授予指定的系统权限策略  |                                          |
| --attach-custom-policy     | 无   | 为该角色授予指定的自定义权限策略 |                                          |
