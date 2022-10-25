# java-sdk

## Usage

1. Import this project.
2. Setup env:

```
# https://www.alibabacloud.com/help/resource-access-management/latest/view-the-basic-information-about-a-ram-role
ALIBABA_CLOUD_ROLE_ARN=<role_arn>

# https://www.alibabacloud.com/help/resource-access-management/latest/manage-an-oidc-idp#section-f4d-9qk-bfl
ALIBABA_CLOUD_OIDC_PROVIDER_ARN=<oidc_provider_arn>

# /path/to/oidc/token/file
ALIBABA_CLOUD_OIDC_TOKEN_FILE=<oidc_token_file>
```

3. Run `Main`:

```
call GetCallerIdentity via oidc token success:

{IdentityType=AssumedRoleUser, AccountId=18***, RequestId=69***, PrincipalId=3***:test-rrsa-oidc-token, UserId=null, Arn=acs:ram::18***:assumed-role/test-rrsa-***/test-rrsa-oidc-token, RoleId=3***}
```