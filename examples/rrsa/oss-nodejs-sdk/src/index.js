const OSS = require("ali-oss");

// @alicloud/credentials >= 2.3.1
const Credential = require("@alicloud/credentials");


function newOidcCredential() {
    // https://www.alibabacloud.com/help/doc-detail/378664.html
    return new Credential.default(new Credential.Config({
      type: 'oidc_role_arn',
      roleArn: process.env.ALIBABA_CLOUD_ROLE_ARN,
      oidcProviderArn: process.env.ALIBABA_CLOUD_OIDC_PROVIDER_ARN,
      oidcTokenFilePath: process.env.ALIBABA_CLOUD_OIDC_TOKEN_FILE,
      roleSessionName: 'test-rrsa-oidc-token',
    }))
}

function newCredential() {
    // https://www.alibabacloud.com/help/doc-detail/378664.html
    return new Credential.default();
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function main() {
    console.log("test oss sdk using rrsa oidc token");
    const durationSeconds = 3600; // 1 hour
    // 两种方法都可以
    const cred = newCredential();
    // or
    // const cred = newOidcCredential();

    const { accessKeyId, accessKeySecret, securityToken } = await cred.getCredential();
    const client = new OSS({
      accessKeyId,
      accessKeySecret,
      stsToken: securityToken,
      refreshSTSTokenInterval: durationSeconds * 0.02 * 1000,
      refreshSTSToken: async () => {
        const { accessKeyId, accessKeySecret, securityToken } = await cred.getCredential();
        return {
          accessKeyId,
          accessKeySecret,
          stsToken: securityToken,
        };
      },
    });

    for (let i = 0; i < 100; i++) {
        console.info(`${i}`)
        await client.listBuckets().then(body => {
            console.log(body.res.status)
            console.log(body.res.headers['x-oss-request-id'])
            console.log("call oss.listBuckets via oidc token success:");
            body.buckets.forEach(item => {
                // console.info(`- ${item.name}`);
            });
        });
        await sleep(1000*durationSeconds)
    }
}

main().catch(console.error);
