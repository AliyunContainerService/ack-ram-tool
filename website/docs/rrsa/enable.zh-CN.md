---
slug: /zh-CN/rrsa/enable
title: enable（中文）
sidebar_position: 1
---

# enable

为特定集群启用 RRSA 特性。

## 使用示例

```shell
$ ack-ram-tool rrsa enable --cluster-id <clusterId>

? Are you sure you want to enable RRSA feature? Yes
2023-04-20T14:30:40+08:00 INFO Enable RRSA feature for cluster c86fdd*** successfully
```

## 命令行参数

```
Usage:
  ack-ram-tool rrsa enable [flags]

Flags:
  -c, --cluster-id string   The cluster id to use
  -h, --help                help for enable

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
