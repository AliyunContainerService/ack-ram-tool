ack-ram-tool
=============================

一个协助您在使用 ACK 过程中便捷执行涉及 RAM、访问凭证、RBAC权限等相关操作的命令行工具。

.. contents::

安装
-----

您可以通过 `Releases <https://github.com/AliyunContainerService/ack-ram-tool/releases>`__ 页面下载最新版的命令行工具。


配置凭证信息
-----------

ack-ram-tool 将通过以下顺序在系统中查找凭证信息：

1. 自动使用环境变量中存在的凭证信息 （
注：程序也支持 `aliyun cli 所支持的凭证相关环境变量 <https://github.com/aliyun/aliyun-cli#support-for-environment-variables>`__ ）:

.. code-block:: shell

    # access key id
    $ export ALIBABA_CLOUD_ACCESS_KEY_ID=foo
    # access key secret
    $ export ALIBABA_CLOUD_ACCESS_KEY_SECRET=bar
    # sts token (可选)
    $ export ALIBABA_CLOUD_SECURITY_TOKEN=foobar

    # or use credentials URI: https://github.com/aliyun/aliyun-cli#use-credentials-uri
    $ export ALIBABA_CLOUD_CREDENTIALS_URI=http://localhost:6666/?user=jacksontian

2. 当环境变量中不存在凭证信息时，如果存在 aliyun cli 的配置文件 ``~/.aliyun/config.json`` (关于 aliyun cli 的配置文件详情请参考
  `官方文档 <https://www.alibabacloud.com/help/doc-detail/110341.htm>`__ ) ，程序将自动使用该配置文件。
3. 当 aliyun cli 的配置文件不存在时，程序将尝试使用 ``~/.alibabacloud/credentials`` 文件中配置的凭证信息（可以通过 ``--profile-file`` 参数指定文件路径）:

.. code-block:: shell

    $ cat ~/.alibabacloud/credentials

    [default]
    type = access_key
    access_key_id = foo
    access_key_secret = bar


使用示例
--------

kubectl/client-go 认证插件
++++++++++++++++++++++++++

一个用于访问 ACK 集群的 `kubectl/client-go 认证插件 <https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins>`__ 。

获取使用该认证插件的 kubeconfig 文件（使用临时 kubeconfig）：

.. code-block:: shell

    ack-ram-tool credential-plugin get-kubeconfig --cluster-id <clusterId> > kubeconfig


使用获取的 kubeconfig 访问集群（在证书过期前会自动获取新的证书）：

.. code-block:: shell

    kubectl --kubeconfig=kubeconfig get ns


清理缓存的访问凭证：

.. code-block:: shell

    rm ~/.kube/cache/ack-ram-tool/credential-plugin/*.json



RAM Roles for Service Accounts (RRSA)
++++++++++++++++++++++++++++++++++++++++

为集群启用 `RRSA 特性 <https://www.alibabacloud.com/help/doc-detail/356611.html>`__ :

.. code-block:: shell

    $ ack-ram-tool rrsa enable -c <clusterId>

    ? Are you sure you want to enable RRSA feature? Yes
    Enable RRSA feature for cluster c86fdd*** successfully


检查当前集群是否已启用 RRSA 特性:

.. code-block:: shell

    $ ack-ram-tool rrsa status -c <clusterId>

    RRSA feature:          enabled
    OIDC Provider Name:    ack-rrsa-c86fdd***
    OIDC Provider Arn:     acs:ram::18***:oidc-provider/ack-rrsa-c86fdd***
    OIDC Token Issuer:     https://oidc-ack-***/c86fdd***


为 RAM 角色关联一个 Service Account（允许使用这个 Service Account 的 OIDC Token 来扮演此 RAM 角色。
通过指定 ``--create-role-if-not-exist`` 参数实现在角色不存在时自动创建对应的 RAM 角色）:

.. code-block:: shell

    $ ack-ram-tool rrsa associate-role -c <clusterId> --create-role-if-not-exist \
        -r <roleName> -n <namespace> -s <serviceAccount>

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


测试使用指定的 OIDC token 扮演 RAM 角色获取 STS Token:

.. code-block:: shell

    $ ack-ram-tool rrsa assume-role -r <roleArn> -p <oidcProviderArn> -t <oidcTokenFile>

    Retrieved a STS token:
    AccessKeyId:       STS.***
    AccessKeySecret:   7UVy***
    SecurityToken:     CAIS***
    Expiration:        2021-12-03T05:51:37Z


可以通过 ``setup-addon`` 命令快速配置集群组件使用 RRSA 特性时所需要的 RAM 相关配置。
比如配置 ``kritis-validation-hook`` 组件所需的 RAM 配置（需要在安装组件前进行配置）:

.. code-block:: shell

    ack-ram-tool rrsa setup-addon --addon-name kritis-validation-hook -c <clusterId>
