---
slug: status
sidebar_position: 4
---

# status

Query the RRSA feature status of a specific cluster.

## Usage

```shell
$ ack-ram-tool rrsa status --cluster-id <clusterId>

RRSA feature:          enabled
OIDC Provider Name:    ack-rrsa-c12d3***
OIDC Provider Arn:     acs:ram::113***:oidc-provider/ack-rrsa-c12d3***
OIDC Token Issuer:     https://oidc-ack-***.aliyuncs.com/c12d3***
```

## Flags

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

Descriptions：

| Flag             | Default | Required | Description |
|------------------|---------|----------|-------------|
| -c, --cluster-id |        | Yes      | Cluster ID  |
