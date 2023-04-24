ack-ram-tool
=============================

一个协助您在使用 ACK 过程中便捷执行涉及 RAM、访问凭证、RBAC权限等相关操作的命令行工具。

.. contents::

安装
-----

您可以通过 `Releases <https://github.com/AliyunContainerService/ack-ram-tool/releases>`__ 页面下载最新版的命令行工具。


配置凭证信息
-----------

详见 `文档 <https://aliyuncontainerservice.github.io/ack-ram-tool/zh-CN/getting-started#%E5%87%AD%E8%AF%81%E4%BF%A1%E6%81%AF>`__

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

    $ ack-ram-tool rrsa enable --cluster-id <clusterId>

    ? Are you sure you want to enable RRSA feature? Yes
    Enable RRSA feature for cluster c86fdd*** successfully

为 RAM 角色关联一个 Service Account（允许使用这个 Service Account 的 OIDC Token 来扮演此 RAM 角色:

.. code-block:: shell

    $ ack-ram-tool rrsa associate-role --cluster-id <clusterId> \
        --namespace <namespce> --service-account <serviceAccountName> \
        --role-name <roleName>

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

文档
--------

更多信息详见 `文档 <https://aliyuncontainerservice.github.io/ack-ram-tool/>`__

