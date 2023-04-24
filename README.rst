ack-ram-tool
=============

A command line utility and library for using RAM、Credential and permission related features in Alibaba Cloud Container Service For Kubernetes (ACK).
`中文文档 <README.zh-cn.rst>`__

.. contents::


Installation
--------------

You can download the latest release from `Releases <https://github.com/AliyunContainerService/ack-ram-tool/releases>`__ page.


Credential
-------------



Usage
--------


kubectl/client-go credential plugin
+++++++++++++++++++++++++++++++++++++

A `kubectl/client-go credential plugin <https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins>`__ for ACK。

Get a kubeconfig with exec credential plugin format：

.. code-block:: shell

    ack-ram-tool credential-plugin get-kubeconfig --cluster-id <clusterId> > kubeconfig


Use this kubeconfig to access cluster:

.. code-block:: shell

    kubectl --kubeconfig=kubeconfig get ns


Remove cached credentials:

.. code-block:: shell

    rm ~/.kube/cache/ack-ram-tool/credential-plugin/*.json



RAM Roles for Service Accounts (RRSA)
++++++++++++++++++++++++++++++++++++++++

Enable `RRSA feature <https://www.alibabacloud.com/help/doc-detail/356611.html>`__ :

.. code-block:: shell

    $ ack-ram-tool rrsa enable --cluster-id <clusterId>

    ? Are you sure you want to enable RRSA feature? Yes
    Enable RRSA feature for cluster c86fdd*** successfully


Associate an RAM Role to a service account (use the ``--create-role-if-not-exist`` flag to
auto create an RAM Role when it doesn't exist):

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

Documentation
---------------


