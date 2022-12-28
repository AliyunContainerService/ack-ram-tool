ack-ram-tool
=============

A command line utility and library for using RAM、Credential and permission related features in Alibaba Cloud Container Service For Kubernetes (ACK).

.. contents::

`中文文档 <README.zh-cn.rst>`__

Installation
--------------

You can download the latest release from `Releases <https://github.com/AliyunContainerService/ack-ram-tool/releases>`__ page.


Credential
-------------

You can reuse ``~/.aliyun/config.json`` file from aliyun cli (For detailed configuration instructions, please visit the document
`Configuration Alibaba Cloud CLI <https://www.alibabacloud.com/help/doc-detail/110341.htm>`__ ).


Or use ``~/.alibabacloud/credentials`` file (this path can be overridden using the ``--profile-file`` flag):

.. code-block:: shell

    $ cat ~/.alibabacloud/credentials

    [default]
    type = access_key
    access_key_id = foo
    access_key_secret = bar

Or environment variables (also support credential related environment variables from `aliyun cli <https://github.com/aliyun/aliyun-cli#support-for-environment-variables>`__):

.. code-block:: shell

    # access key id
    $ export ALIBABA_CLOUD_ACCESS_KEY_ID=foo
    # access key secret
    $ export ALIBABA_CLOUD_ACCESS_KEY_SECRET=bar
    # sts token (optional)
    $ export ALIBABA_CLOUD_SECURITY_TOKEN=foobar

    # or use credentials URI: https://github.com/aliyun/aliyun-cli#use-credentials-uri
    $ export ALIBABA_CLOUD_CREDENTIALS_URI=http://localhost:6666/?user=jacksontian


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

    $ ack-ram-tool rrsa enable -c <clusterId>

    ? Are you sure you want to enable RRSA feature? Yes
    Enable RRSA feature for cluster c86fdd*** successfully



Check status of RRSA feature:

.. code-block:: shell

    $ ack-ram-tool rrsa status -c <clusterId>

    RRSA feature:          enabled
    OIDC Provider Name:    ack-rrsa-c86fdd***
    OIDC Provider Arn:     acs:ram::18***:oidc-provider/ack-rrsa-c86fdd***
    OIDC Token Issuer:     https://oidc-ack-***/c86fdd***


Associate an RAM Role to a service account (use the ``--create-role-if-not-exist`` flag to
auto create an RAM Role when it doesn't exist):

.. code-block:: shell

    $ ack-ram-tool rrsa associate-role --create-role-if-not-exist -c <clusterId> -r <roleName> -n <namespace> -s <serviceAccount>

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


Testing assume role with give OIDC token:

.. code-block:: shell

    $ ack-ram-tool rrsa assume-role -r <roleArn> -p <oidcProviderArn> -t <oidcTokenFile>

    Retrieved a STS token:
    AccessKeyId:       STS.***
    AccessKeySecret:   7UVy***
    SecurityToken:     CAIS***
    Expiration:        2021-12-03T05:51:37Z


The `setup-addon` command allows you to quickly configure the RAM-related configuration
required for the cluster components to use the RRSA feature.
For example, configure the RAM configuration required for the `kritis-validation-hook` 
component (needs to be configured before installing the component):

.. code-block:: shell

    ack-ram-tool rrsa setup-addon --addon-name kritis-validation-hook -c <clusterId>


Disable RRSA feature:

.. code-block:: shell

    $ ack-ram-tool rrsa disable -c <clusterId>

    ? Are you sure you want to disable RRSA feature? Yes
    Disable RRSA feature for cluster c86fdd*** successfully

