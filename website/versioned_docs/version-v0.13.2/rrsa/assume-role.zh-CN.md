---
slug: /zh-CN/rrsa/assume-role
title: assume-role（中文）
sidebar_position: 5
---

# assume-role

测试使用 oidc token 扮演特定 RAM 角色。

## 使用示例

```shell
$ ack-ram-tool rrsa assume-role --oidc-provider-arn <oidcProviderArn> \
  --role-arn <roleArn> --oidc-token-file <pathToTokenFile>

    Retrieved a STS token:
    AccessKeyId:       STS.***
    AccessKeySecret:   7UVy***
    SecurityToken:     CAIS***
    Expiration:        2021-12-03T05:51:37Z

```

## 命令行参数

```
Usage:
  ack-ram-tool rrsa assume-role [flags]

Flags:
  -h, --help                       help for assume-role
  -p, --oidc-provider-arn string   The arn of OIDC provider
  -t, --oidc-token-file string     Path to OIDC token file. If value is '-', will read token from stdin
  -r, --role-arn string            The arn of RAM role

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

| 参数名称                       | 默认值 | 必需参数 | 说明                                                         |
|----------------------------|-----|------|------------------------------------------------------------|
|-p, --oidc-provider-arn| 无   | 是    | 为集群注册的 RAM 角色 SSO 供应商 ARN                                  |
|-r, --role-arn| 无   | 是    | 被扮演的 RAM 角色的 ARN                                           |
|-t, --oidc-token-file| 无   | 是    | oidc token 文件的路径。当值为 `-` 时支持从标准输入从读取 token（比如通过管道传递 token） |
