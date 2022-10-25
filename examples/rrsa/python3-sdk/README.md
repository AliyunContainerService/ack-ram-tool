# python3-sdk

## Usage

```bash
$ python3 -m venv venv
$ . ./venv/bin/activate
$ pip install -r requirements.txt

# https://www.alibabacloud.com/help/resource-access-management/latest/view-the-basic-information-about-a-ram-role
$ export ALIBABA_CLOUD_ROLE_ARN=<role_arn>

# https://www.alibabacloud.com/help/resource-access-management/latest/manage-an-oidc-idp#section-f4d-9qk-bfl
$ export ALIBABA_CLOUD_OIDC_PROVIDER_ARN=<oidc_provider_arn>

# /path/to/oidc/token/file
$ export ALIBABA_CLOUD_OIDC_TOKEN_FILE=<oidc_token_file>

$ python main.py

test open api sdk use rrsa oidc token
call sts.GetCallerIdentity via oidc token success:
{
  "headers": {
    ...
  },
  "body": {
    "AccountId": "18***",
    "Arn": "***",
    "IdentityType": "AssumedRoleUser",
    "PrincipalId": "***:auth-with-rrsa-oidc-token",
    "RequestId": "4FCA5B69-***",
    "RoleId": "***"
  }
}

```
