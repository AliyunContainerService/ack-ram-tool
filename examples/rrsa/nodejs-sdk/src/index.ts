import Sts20150401, * as $Sts20150401 from '@alicloud/sts20150401';
import OpenApi, * as $OpenApi from '@alicloud/openapi-client';
import * as $tea from '@alicloud/tea-typescript';
import Credential, { Config } from '@alicloud/credentials';


export default class Client {

  static createClient(): Sts20150401 {
    let cred = new Credential(new Config({
      type:                   'oidc_role_arn',
      roleArn:                process.env.ALIBABA_CLOUD_ROLE_ARN,
      oidcProviderArn:        process.env.ALIBABA_CLOUD_OIDC_PROVIDER_ARN,
      oidcTokenFilePath:      process.env.ALIBABA_CLOUD_OIDC_TOKEN_FILE,
      roleSessionName:        'test-rrsa-oidc-token',
    }));
    let config = new $OpenApi.Config({
      credential: cred,
    });
    // get endpoint from https://www.alibabacloud.com/help/resource-access-management/latest/endpoints
    config.endpoint = `sts.aliyuncs.com`;
    return new Sts20150401(config);
  }

  static async main(args: string[]): Promise<void> {
    let client = Client.createClient();
    await client.getCallerIdentity().then(body =>  {
      console.info(body)
    });
  }

}

Client.main(process.argv.slice(2));
