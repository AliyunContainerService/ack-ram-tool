---
slug: /zh-CN/rrsa/install-helper-addon
title: install-helper-addon（中文）
sidebar_position: 3
---

# install-helper-addon

在集群内安装 RRSA 辅助组件 [ack-pod-identity-webhook](https://www.alibabacloud.com/help/doc-detail/600451.html)。

## 使用示例

```shell
$ ack-ram-tool rrsa install-helper-addon --cluster-id <clusterId>

? Are you sure you want to install ack-pod-identity-webhook? Yes
2023-04-20T15:39:41+08:00 INFO Start to install ack-pod-identity-webhook
2023-04-20T15:40:49+08:00 INFO Install ack-pod-identity-webhook for cluster c12d3*** successfully
```

## 命令行参数

```
Usage:
  ack-ram-tool rrsa install-helper-addon [flags]

Flags:
  -c, --cluster-id string   The cluster id to use
  -h, --help                help for install-helper-addon

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
