---
slug: /
sidebar_position: 1
---

# Getting started


## Installation

lease go to the [Releases](https://github.com/AliyunContainerService/ack-ram-tool/releases) page
to download the latest version of the ack-ram-tool.


## Configuration


### Credentials

ack-ram-tool will search for credential information in the system in the following order：

1. Automatically use credential information that exists in the environment variables（
   Note: This tool also supports the credential-related environment variables supported by [aliyun cli](https://github.com/aliyun/aliyun-cli#support-for-environment-variables) ）:

```
# access key id
$ export ALIBABA_CLOUD_ACCESS_KEY_ID=foo

# access key secret
$ export ALIBABA_CLOUD_ACCESS_KEY_SECRET=bar

# sts token (Optional)
$ export ALIBABA_CLOUD_SECURITY_TOKEN=foobar

# or use credentials URI: https://github.com/aliyun/aliyun-cli#use-credentials-uri
$ export ALIBABA_CLOUD_CREDENTIALS_URI=http://localhost:6666/?user=jacksontian
```

2. When credential information does not exist in the environment variables, if there is an aliyun cli configuration file
   ``~/.aliyun/config.json`` (For details on the aliyun cli configuration file, 
   please refer to the [official documentation](https://www.alibabacloud.com/help/doc-detail/110341.htm) ) , 
   the program will automatically use that configuration file.

3. When the aliyun cli configuration file does not exist, the program will attempt to use the credential information
  configured in the ``~/.alibabacloud/credentials`` file (which can be specified by the ``--profile-file`` flags):

```
$ cat ~/.alibabacloud/credentials

[default]
type = access_key
access_key_id = foo
access_key_secret = bar
```


## Permissions

In order to use ack-ram-tool normally, you need to grant the necessary RAM permissions and RBAC permissions 
for the Alibaba Cloud RAM user or RAM role that uses this tool. 
For the minimum permission information required for each subcommand, please refer to [Permissions](permissions).
