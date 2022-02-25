# nodejs-sdk

## Usage

```bash
$ export ALIBABA_CLOUD_ROLE_ARN=<role_arn>
$ export ALIBABA_CLOUD_OIDC_PROVIDER_ARN=<oidc_provider_arn>
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
