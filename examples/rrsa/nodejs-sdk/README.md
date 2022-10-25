# nodejs-sdk

## Usage

```bash
# https://www.alibabacloud.com/help/resource-access-management/latest/view-the-basic-information-about-a-ram-role
$ export ALIBABA_CLOUD_ROLE_ARN=<role_arn>

# https://www.alibabacloud.com/help/resource-access-management/latest/manage-an-oidc-idp#section-f4d-9qk-bfl
$ export ALIBABA_CLOUD_OIDC_PROVIDER_ARN=<oidc_provider_arn>

# /path/to/oidc/token/file
$ export ALIBABA_CLOUD_OIDC_TOKEN_FILE=<oidc_token_file>

$ npm run demo

> nodejs-sdk@1.0.0 demo /.../ack-ram-tool/examples/rrsa/nodejs-sdk
> ts-node src/index.ts

GetCallerIdentityResponse {
  headers:
   { ... },
  body:
   GetCallerIdentityResponseBody {
     identityType: 'AssumedRoleUser',
     accountId: '18***',
     requestId: 'F4***',
     principalId: '3***:test-rrsa-oidc-token',
     arn:
      'acs:ram::18***:assumed-role/test-rrsa-***/test-rrsa-oidc-token',
     roleId: '3***' } }

```
