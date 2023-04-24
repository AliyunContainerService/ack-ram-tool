---
slug: /zh-CN/export-credentials/export-credentials
title: export-credentials（中文）
---

# export-credentials

导出获取到的凭证信息或者使用该凭证执行外部程序。

## 使用示例

### 默认输出

```shell
$ ack-ram-tool export-credentials
{
  "mode": "AK",
  "access_key_id": "LT***",
  "access_key_secret": "vHLE***"
 }
```

### --format aliyun-cli-uri-json

```shell
$ ack-ram-tool export-credentials --format aliyun-cli-uri-json
{
  "Code": "Success",
  "AccessKeyId": "LT***",
  "AccessKeySecret": "vHLE***",
  "SecurityToken": "",
  "Expiration": "2023-04-20T12:09:37Z"
 }
```

### --format ecs-metadata-json

```shell
$ ack-ram-tool export-credentials --format ecs-metadata-json
{
  "Code": "Success",
  "AccessKeyId": "LT***",
  "AccessKeySecret": "vHLE***",
  "SecurityToken": "",
  "Expiration": "2023-04-20T12:11:04Z"
 }
```

### --format credential-file-ini

```shell
$ ack-ram-tool export-credentials --format credential-file-ini
[default]
enable = true
type = access_key
access_key_id = LT***
access_key_secret = vHLE***
```

### --format environment-variables

```shell
$ ack-ram-tool export-credentials --format environment-variables

for aliyun cli:

export ALIBABACLOUD_ACCESS_KEY_ID=LT***
export ALIBABACLOUD_ACCESS_KEY_SECRET=vHLE***

for terraform:

export ALICLOUD_ACCESS_KEY=LT***
export ALICLOUD_SECRET_KEY=vHLE***

for other tools:

export ALIBABA_CLOUD_ACCESS_KEY_ID=LT***
export ALICLOUD_ACCESS_KEY=LT***
export ALIBABACLOUD_ACCESS_KEY_ID=LT***
export ALICLOUD_SECRET_KEY=LT***
export ALIBABA_CLOUD_ACCESS_KEY_SECRET=vHLE***
export ALIBABACLOUD_ACCESS_KEY_SECRET=vHLE***
```

### --format aliyun-cli-uri-json --serve ADDR

```shell
$ ack-ram-tool export-credentials --format aliyun-cli-uri-json --serve 127.0.0.1:1234
2023-04-20T20:05:40+08:00 WARN Serving HTTP on 127.0.0.1:1234
$ curl http://127.0.0.1:1234
{
  "Code": "Success",
  "AccessKeyId": "LT***",
  "AccessKeySecret": "vHLE***",
  "SecurityToken": "",
  "Expiration": "2023-04-20T12:14:15Z"
 }
```

### --format aliyun-cli-uri-json -- COMMAND [ARGS]

```shell
$ ack-ram-tool export-credentials --format environment-variables -- aliyun sts GetCallerIdentity
{
	"AccountId": "113***",
	"Arn": "acs:ram::113***:user/***",
	"IdentityType": "RAMUser",
	"PrincipalId": "272***",
	"RequestId": "28B93***",
	"UserId": "272***"
}
```

## 命令行参数

```
Usage:
  ack-ram-tool export-credentials [flags]

Flags:
  -f, --format string   The output format to display credentials (aliyun-cli-config-json, aliyun-cli-uri-json, ecs-metadata-json, credential-file-ini, environment-variables) (default "aliyun-cli-config-json")
  -h, --help            help for export-credentials
  -s, --serve string    start a server to export credentials

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

| 参数名称         | 默认值                    | 必需参数 | 说明                                          |
|--------------|------------------------|------|---------------------------------------------|
| -f, --format | aliyun-cli-config-json | 否    | 指定输出格式，详见使用示例                               |
| -s, --serve  | 无                      | 否    | 启动一个监听指定地址的 HTTP Server，访问该服务将返回凭证信息，详见使用示例 |
