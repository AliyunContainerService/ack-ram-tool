# coding: utf-8
import json
import os

import oss2

# alibabacloud-credentials>=0.3.0
from alibabacloud_credentials.client import Client as CredClient
from alibabacloud_credentials.models import Config as CredConfig

from alibabacloud_sts20150401.client import Client as Sts20150401Client
from alibabacloud_tea_openapi import models as sts_open_api_models
from alibabacloud_sts20150401 import models as sts_20150401_models
from alibabacloud_tea_util import models as util_models
from alibabacloud_tea_util.client import Client as UtilClient

from oss2.credentials import (
    CredentialsProvider as OSSCredentialsProvider,
    Credentials as OSSCredentials
)

ENV_ROLE_ARN = "ALIBABA_CLOUD_ROLE_ARN"
ENV_OIDC_PROVIDER_ARN = "ALIBABA_CLOUD_OIDC_PROVIDER_ARN"
ENV_OIDC_TOKEN_FILE = "ALIBABA_CLOUD_OIDC_TOKEN_FILE"


def test_open_api_sdk(cred):
    config = sts_open_api_models.Config(
        # get endpoint from https://www.alibabacloud.com/help/resource-access-management/latest/endpoints
        endpoint="sts.aliyuncs.com", credential=cred)
    client = Sts20150401Client(config)
    resp = client.get_caller_identity()
    print("call sts.GetCallerIdentity via oidc token success:\n{}\n".format(
        json.dumps(resp.to_map(), indent=2)))


def test_oss_sdk(cred):
    endpoint = 'http://oss-cn-hangzhou.aliyuncs.com'
    provider = OSSOidcCredentialProvider(cred)
    auth = oss2.ProviderAuth(provider)
    service = oss2.Service(auth=auth, endpoint=endpoint)
    resp = service.list_buckets()
    print("call oss.listBuckets via oidc token success:")
    for bucket in resp.buckets:
        print('- {}'.format(bucket.name))


def get_oidc_cred(role_arn, oidc_arn, token_file):
    config = CredConfig(type='oidc_role_arn', role_arn=role_arn,
                        oidc_provider_arn=oidc_arn, oidc_token_file_path=token_file,
                        role_session_name='auth-with-rrsa-oidc-token')
    cred = CredClient(config)
    return cred


class OSSOidcCredentialProvider(OSSCredentialsProvider):
    def __init__(self, cred):
        self._cred = cred

    def get_credentials(self):
        access_key_id = self._cred.get_access_key_id()
        access_key_secret = self._cred.get_access_key_secret()
        security_token = self._cred.get_security_token()
        return OSSCredentials(access_key_id=access_key_id, access_key_secret=access_key_secret,
                              security_token=security_token)


def main():
    cred = get_oidc_cred(os.environ[ENV_ROLE_ARN], os.environ[ENV_OIDC_PROVIDER_ARN],
                         os.environ[ENV_OIDC_TOKEN_FILE])

    # test open api sdk (https://github.com/aliyun/alibabacloud-python-sdk) use rrsa oidc token
    print("\ntest open api sdk use rrsa oidc token")
    test_open_api_sdk(cred)

    # test oss sdk (https://github.com/aliyun/aliyun-oss-python-sdk) use rrsa oidc token
    if os.getenv("TEST_OSS_SDK") == "true":
        print("\ntest oss sdk use rrsa oidc token")
        test_oss_sdk(cred)


if __name__ == '__main__':
    main()
