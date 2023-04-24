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

```
# access key id
$ export ALIBABA_CLOUD_ACCESS_KEY_ID=foo

# access key secret
$ export ALIBABA_CLOUD_ACCESS_KEY_SECRET=bar

# sts token (可选)
$ export ALIBABA_CLOUD_SECURITY_TOKEN=foobar

# or use credentials URI: https://github.com/aliyun/aliyun-cli#use-credentials-uri
$ export ALIBABA_CLOUD_CREDENTIALS_URI=http://localhost:6666/?user=jacksontian
```

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
