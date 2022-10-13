# python3-sdk

## Usage

```bash
$ python3 -m venv venv
$ . ./venv/bin/activate
$ pip install -r requirements.txt

$ export ALIBABA_CLOUD_ROLE_ARN=<role_arn>
$ export ALIBABA_CLOUD_OIDC_PROVIDER_ARN=<oidc_provider_arn>
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
