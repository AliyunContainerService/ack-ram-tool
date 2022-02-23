# java-sdk

## Usage

1. Import this project.
2. Setup env:

```
ALIBABA_CLOUD_ROLE_ARN=<role_arn>
ALIBABA_CLOUD_OIDC_PROVIDER_ARN=<oidc_provider_arn>
ALIBABA_CLOUD_OIDC_TOKEN_FILE=<oidc_token_file>
```

3. Run `Main`:

```
{IdentityType=AssumedRoleUser, AccountId=18***, RequestId=69***, PrincipalId=3***:test-rrsa-oidc-token, UserId=null, Arn=acs:ram::18***:assumed-role/test-rrsa-***/test-rrsa-oidc-token, RoleId=3***}
```