# coding: utf-8
import os

# alibabacloud-credentials>=0.3.1
from alibabacloud_credentials.client import Client as CredClient
from alibabacloud_credentials.models import Config as CredConfig

from alibabacloud_cs20151215.client import Client as CS20151215Client
from alibabacloud_cs20151215 import models as csmodules
from alibabacloud_tea_openapi import models as open_api_models

import oss2
from oss2.credentials import (
    CredentialsProvider as OSSCredentialsProvider,
    Credentials as OSSCredentials
)

ENV_ROLE_ARN = "ALIBABA_CLOUD_ROLE_ARN"
ENV_OIDC_PROVIDER_ARN = "ALIBABA_CLOUD_OIDC_PROVIDER_ARN"
ENV_OIDC_TOKEN_FILE = "ALIBABA_CLOUD_OIDC_TOKEN_FILE"


def test_open_api_sdk(cred):
    config = open_api_models.Config(credential=cred)
    config.endpoint = 'cs.cn-hangzhou.aliyuncs.com'

    client = CS20151215Client(config)
    resp = client.describe_clusters(csmodules.DescribeClustersRequest())

    print("call cs.describe_clusters via oidc token success:\n")
    for cluster in resp.body:
        print("cluster id: {}, cluster name: {}".format(cluster.cluster_id, cluster.name))
    print('\n')


def new_cred():
    # https://www.alibabacloud.com/help/doc-detail/2567976.html
    cred = CredClient()
    return cred


def new_oidc_cred():
    # https://www.alibabacloud.com/help/doc-detail/2567976.html
    config = CredConfig(
        type='oidc_role_arn',
        role_arn=os.environ[ENV_ROLE_ARN],
        oidc_provider_arn=os.environ[ENV_OIDC_PROVIDER_ARN],
        oidc_token_file_path=os.environ[ENV_OIDC_TOKEN_FILE],
        role_session_name='auth-with-rrsa-oidc-token')

    # https://next.api.aliyun.com/product/Sts
    # config.sts_endpoint = 'sts-vpc.cn-hangzhou.aliyuncs.com'

    cred = CredClient(config)
    return cred


def main():
    # 两种方法都可以
    cred = new_cred()
    # or
    # cred = new_oidc_cred()

    # test open api sdk (https://github.com/aliyun/alibabacloud-python-sdk) use rrsa oidc token
    print("\ntest open api sdk use rrsa oidc token")
    test_open_api_sdk(cred)


if __name__ == '__main__':
    main()
