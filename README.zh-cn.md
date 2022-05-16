# ack-ram-tool

一个辅助配置在使用 ACK 过程中涉及的 RAM 相关操作的命令行工具及 Golang 库。


## 安装

您可以通过 [Releases](https://github.com/AliyunContainerService/ack-ram-tool/releases) 页面下载最新版的命令行工具。


## 配置认证信息

您可以通过 `~/.alibabacloud/credentials` 文件配置认证信息（也可以通过 `--profile-file` 参数指定文件路径）:

```
$ cat ~/.alibabacloud/credentials

[default]
type = access_key
access_key_id = foo
access_key_secret = bar
```

您也可以通过环境变量配置认证信息:

```
$ export ALIBABA_CLOUD_ACCESS_KEY_ID=foo
$ export ALIBABA_CLOUD_ACCESS_KEY_SECRET=bar
```

## 使用示例


### RAM Roles for Service Accounts (RRSA)

为集群启用 [RRSA 特性](https://www.alibabacloud.com/help/doc-detail/356611.html):

```
$ ack-ram-tool rrsa enable -c <clusterId>

? Are you sure you want to enable RRSA feature? Yes
Enable RRSA feature for cluster c86fdd*** successfully

```


检查当前集群是否已启用 RRSA 特性:

```
$ ack-ram-tool rrsa status -c <clusterId>

RRSA feature:          enabled
OIDC Provider Name:    ack-rrsa-c86fdd***
OIDC Provider Arn:     acs:ram::18***:oidc-provider/ack-rrsa-c86fdd***
OIDC Token Issuer:     https://oidc-ack-***/c86fdd***

```

禁用 RRSA 特性:

```
$ ack-ram-tool rrsa disable -c <clusterId>

? Are you sure you want to disable RRSA feature? Yes
Disable RRSA feature for cluster c86fdd*** successfully

```

为 RAM 角色关联一个 Service Account（允许使用这个 Service Account 的 OIDC Token 来扮演此 RAM 角色。
通过指定 ``--create-role-if-not-exist`` 参数实现在角色不存在时自动创建对应的 RAM 角色）:

```
$ ack-ram-tool rrsa associate-role -c <clusterId> --create-role-if-not-exist -r <roleName> -n <namespace> -s <serviceAccount>

? Are you sure you want to associate RAM Role test-rrsa to service account test-serviceaccount (namespace: test-namespace)? Yes
Will change the assumeRolePolicyDocument of RAM Role test-rrsa with blow content:
{
  "Statement": [
   {
    "Action": "sts:AssumeRole",
    "Effect": "Allow",
    "Principal": {
     "RAM": [
      "acs:ram::18***:root"
     ]
    }
   },
   },
   {
    "Action": "sts:AssumeRole",
    "Condition": {
     "StringEquals": {
      "oidc:aud": "sts.aliyuncs.com",
      "oidc:iss": "https://oidc-ack-**/c86fdd***",
      "oidc:sub": "system:serviceaccount:test-namespace:test-serviceaccount"
     }
    },
    "Effect": "Allow",
    "Principal": {
     "Federated": [
      "acs:ram::18***:oidc-provider/ack-rrsa-c86fdd***"
     ]
    }
   }
  ],
  "Version": "1"
 }
? Are you sure you want to associate RAM Role test-rrsa to service account test-serviceaccount (namespace: test-namespace)? Yes
Associate RAM Role test-rrsa to service account test-serviceaccount (namespace: test-namespace) successfully

```

测试使用指定的 OIDC token 扮演 RAM 角色获取 STS Token:

```
$ ack-ram-tool rrsa assume-role -r <roleArn> -p <oidcProviderArn> -t <oidcTokenFile>

Retrieved a STS token:
AccessKeyId:       STS.***
AccessKeySecret:   7UVy***
SecurityToken:     CAIS***
Expiration:        2021-12-03T05:51:37Z

```
