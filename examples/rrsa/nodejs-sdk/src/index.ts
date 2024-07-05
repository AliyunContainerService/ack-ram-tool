import CS20151215, * as $CS20151215 from '@alicloud/cs20151215';
import * as $OpenApi from '@alicloud/openapi-client';

// @alicloud/credentials >= 2.2.6
import Credential, {Config} from '@alicloud/credentials';


export default class Demo {

  static newOidcCredential(): Credential {
    // https://www.alibabacloud.com/help/doc-detail/378664.html
    return new Credential(new Config({
      type: 'oidc_role_arn',
      roleArn: process.env.ALIBABA_CLOUD_ROLE_ARN,
      oidcProviderArn: process.env.ALIBABA_CLOUD_OIDC_PROVIDER_ARN,
      oidcTokenFilePath: process.env.ALIBABA_CLOUD_OIDC_TOKEN_FILE,
      roleSessionName: 'test-rrsa-oidc-token',
    }))
  }

  static newCredential(): Credential {
    // https://www.alibabacloud.com/help/doc-detail/378664.html
    return new Credential();
  }

  static async main(args: string[]): Promise<void> {
    // 两种方法都可以
    const cred = Demo.newCredential();
    // or
    // const cred = Demo.newOidcCredential();

    const config =  new $OpenApi.Config({credential: cred})
    config.endpoint = 'cs.cn-hangzhou.aliyuncs.com';
    const client = new CS20151215(config);
    const req = new $CS20151215.DescribeClustersRequest({})
    await client.describeClusters(req).then(body =>  {
      body.body.forEach((value) => {
        console.info(`cluster id: ${value.clusterId}, cluster name: ${value.name}`)
      })
    });
  }

}

Demo.main(process.argv.slice(2));
