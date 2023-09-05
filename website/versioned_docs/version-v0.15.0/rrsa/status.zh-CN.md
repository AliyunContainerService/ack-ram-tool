---
slug: /zh-CN/rrsa/status
title: status（中文）
sidebar_position: 4
---

# status

查询特定集群的 RRSA 特性状态。

## 使用示例

```shell
$ ack-ram-tool rrsa status --cluster-id <clusterId>

RRSA feature:          enabled
OIDC Provider Name:    ack-rrsa-c12d3***
OIDC Provider Arn:     acs:ram::113***:oidc-provider/ack-rrsa-c12d3***
OIDC Token Issuer:     https://oidc-ack-***.aliyuncs.com/c12d3***
```

## 命令行参数

```
Usage:
  ack-ram-tool rrsa status [flags]

Flags:
  -c, --cluster-id string   The cluster id to use
  -h, --help                help for status

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

| 参数名称             | 默认值 | 必需参数 | 说明    |
|------------------|-----|------|-------|
| -c, --cluster-id | 无   | 是    | 集群 ID |
