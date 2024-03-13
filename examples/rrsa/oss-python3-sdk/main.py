# coding: utf-8
import os

# alibabacloud-credentials>=0.3.1
from alibabacloud_credentials.client import Client as CredClient
from alibabacloud_credentials.models import Config as CredConfig

import oss2
from oss2.credentials import (
    CredentialsProvider as OSSCredentialsProvider,
    Credentials as OSSCredentials
)

ENV_ROLE_ARN = "ALIBABA_CLOUD_ROLE_ARN"
ENV_OIDC_PROVIDER_ARN = "ALIBABA_CLOUD_OIDC_PROVIDER_ARN"
ENV_OIDC_TOKEN_FILE = "ALIBABA_CLOUD_OIDC_TOKEN_FILE"


def test_oss_sdk(cred):
    endpoint = 'http://oss-cn-hangzhou.aliyuncs.com'
    provider = OSSOidcCredentialProvider(cred)
    auth = oss2.ProviderAuth(provider)

    service = oss2.Service(auth=auth, endpoint=endpoint)
    resp = service.list_buckets()

    print("call oss.listBuckets via oidc token success:")
    for bucket in resp.buckets:
        print('- {}'.format(bucket.name))


def new_cred():
    # https://www.alibabacloud.com/help/doc-detail/378661.html
    cred = CredClient()
    return cred


def new_oidc_cred():
    # https://www.alibabacloud.com/help/doc-detail/378661.html
    config = CredConfig(
        type='oidc_role_arn',
        role_arn=os.environ[ENV_ROLE_ARN],
        oidc_provider_arn=os.environ[ENV_OIDC_PROVIDER_ARN],
        oidc_token_file_path=os.environ[ENV_OIDC_TOKEN_FILE],
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
    # 两种方法都可以
    cred = new_cred()
    # or
    # cred = new_oidc_cred()

    # test oss sdk (https://github.com/aliyun/aliyun-oss-python-sdk) use rrsa oidc token
    print("\ntest oss sdk use rrsa oidc token")
    test_oss_sdk(cred)


if __name__ == '__main__':
    main()
