---
slug: assume-role
sidebar_position: 5
---

# assume-role

Test using an OIDC token to assume a specific RAM role.

## Usage

```shell
$ ack-ram-tool rrsa assume-role --oidc-provider-arn <oidcProviderArn> \
  --role-arn <roleArn> --oidc-token-file <pathToTokenFile>

    Retrieved a STS token:
    AccessKeyId:       STS.***
    AccessKeySecret:   7UVy***
    SecurityToken:     CAIS***
    Expiration:        2021-12-03T05:51:37Z

```

## Flags

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

Descriptions：

| Flag                    | Default | Required | Description                                                                                                                                                        |
|-------------------------|---------|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| -p, --oidc-provider-arn |        | Yes      | OIDC Provider ARN                                                                                                                                                  |
| -r, --role-arn          |        | Yes      | Role ARN                                                                                                                                                           |
| -t, --oidc-token-file   |        | Yes      | The path to the OIDC token file. If the value is "-", it mean that the token can be read from standard input(for example, by passing the token through a pipeline) |
