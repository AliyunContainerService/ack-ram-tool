# go-sdk


## Usage

```bash
# https://www.alibabacloud.com/help/resource-access-management/latest/view-the-basic-information-about-a-ram-role
$ export ALIBABA_CLOUD_ROLE_ARN=<role_arn>

# https://www.alibabacloud.com/help/resource-access-management/latest/manage-an-oidc-idp#section-f4d-9qk-bfl
$ export ALIBABA_CLOUD_OIDC_PROVIDER_ARN=<oidc_provider_arn>

# /path/to/oidc/token/file
$ export ALIBABA_CLOUD_OIDC_TOKEN_FILE=<oidc_token_file>

$ go run main.go

call GetCallerIdentity via oidc token success:
{
   "headers": {
      ...
   },
   "body": {
      "IdentityType": "AssumedRoleUser",
      "AccountId": "18***",
      "RequestId": "ABC***",
      "PrincipalId": "3***:test-rrsa-oidc-token",
      "Arn": "acs:ram::18***:assumed-role/test-rrsa-***/test-rrsa-oidc-token",
      "RoleId": "3***"
   }
}

```
