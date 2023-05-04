---
slug: /zh-CN/getting-started
sidebar_position: 1
---

# 新手指南


## 安装

请前往 [Releases](https://github.com/AliyunContainerService/ack-ram-tool/releases) 页面下载最新版的命令行工具。


## 配置

### 凭证信息


ack-ram-tool 将按照以下顺序在系统中查找凭证信息：

1. 自动使用环境变量中存在的凭证信息 （
   注：本工具也支持aliyun cli 所支持的凭证相关[环境变量](https://github.com/aliyun/aliyun-cli#support-for-environment-variables) ）:

| 环境变量                                                                                                                                                                       | 含义                                                                          |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| `ALIBABA_CLOUD_ACCESS_KEY_ID`、`ALICLOUD_ACCESS_KEY`、`ALIBABACLOUD_ACCESS_KEY_ID`、`ALICLOUD_ACCESS_KEY_ID`、`ALIBABACLOUD_ACCESS_KEY_ID`、`ACCESS_KEY_ID`                     | access key id                                                               |
| `ALIBABA_CLOUD_ACCESS_KEY_SECRET`、`ALICLOUD_SECRET_KEY`、`ALIBABACLOUD_ACCESS_KEY_SECRET`、`ALICLOUD_ACCESS_KEY_SECRET`、`ALIBABACLOUD_ACCESS_KEY_SECRET`、`ACCESS_KEY_SECRET` | access key secret                                                           |
| `ALIBABA_CLOUD_SECURITY_TOKEN`、`ALICLOUD_ACCESS_KEY_STS_TOKEN`、`ALIBABACLOUD_SECURITY_TOKEN`、`ALICLOUD_SECURITY_TOKEN`、`ALIBABACLOUD_SECURITY_TOKEN`、`SECURITY_TOKEN`      | sts token                                                                   |
| `ALIBABA_CLOUD_CREDENTIALS_URI`                                                                                                                                            | [credentials URI](https://github.com/aliyun/aliyun-cli#use-credentials-uri) |
| `ALIBABA_CLOUD_ROLE_ARN`                                                                                                                                                   | RAM Role ARN                                                                |
| `ALIBABA_CLOUD_OIDC_PROVIDER_ARN`                                                                                                                                          | OIDC Provider ARN                                                           |
| `ALIBABA_CLOUD_OIDC_TOKEN_FILE`                                                                                                                                            | OIDC Token File                                                             |


2. 当环境变量中不存在凭证信息时，如果存在 aliyun cli 的配置文件 ``~/.aliyun/config.json`` (关于 aliyun cli 的配置文件详情请参考
   [官方文档](https://www.alibabacloud.com/help/doc-detail/110341.htm) ) ，程序将自动使用该配置文件。

3. 当 aliyun cli 的配置文件不存在时，程序将尝试使用 ``~/.alibabacloud/credentials`` 文件中配置的凭证信息（可以通过 ``--profile-file`` 参数指定文件路径）:

```
$ cat ~/.alibabacloud/credentials

[default]
type = access_key
access_key_id = foo
access_key_secret = bar
```


## 权限

为了正常使用 ack-ram-tool，您需要为使用改工具的阿里云 RAM 用户或 RAM 角色授予所需的 RAM 权限和 RBAC 权限。
各个子命令所需的最小权限信息详见 [权限](permissions)
