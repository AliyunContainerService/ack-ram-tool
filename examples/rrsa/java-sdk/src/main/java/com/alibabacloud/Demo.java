package com.alibabacloud;

// com.aliyun:credentials-java >= 0.2.10
import com.aliyun.credentials.Client;
import com.aliyun.cs20151215.models.DescribeClustersRequest;
import com.aliyun.cs20151215.models.DescribeClustersResponse;

import java.util.List;

class TestOpenAPISDK {

    public void CallAPI(com.aliyun.credentials.Client cred) throws Exception {
        com.aliyun.teaopenapi.models.Config config = new com.aliyun.teaopenapi.models.Config();
        config.credential = cred;
        config.endpoint =  "cs.cn-hangzhou.aliyuncs.com";

        // init client
        com.aliyun.cs20151215.Client client = new com.aliyun.cs20151215.Client(config);

        // call open api
        DescribeClustersResponse resp = client.describeClusters(new DescribeClustersRequest());
        System.out.println("call cs.describeClusters via oidc token success:\n");
        List<DescribeClustersResponse.DescribeClustersResponseBody> body = resp.getBody();
        for (DescribeClustersResponse.DescribeClustersResponseBody value : body) {
            System.out.printf("cluster id: %s, cluster name: %s\n", value.clusterId, value.name);
        }
        System.out.println();
    }
}

public class Demo {

    public static void main(String[] args) throws Exception {
        // 两种方式都可以
        com.aliyun.credentials.Client cred = new com.aliyun.credentials.Client();
        // or
        // com.aliyun.credentials.Client cred = newOidcCred();


        // test open api sdk (https://github.com/aliyun/alibabacloud-java-sdk) use rrsa oidc token
        System.out.println("\n");
        System.out.println("test open api sdk use rrsa oidc token");
        TestOpenAPISDK openapiSdk = new TestOpenAPISDK();
        openapiSdk.CallAPI(cred);

    }

    static com.aliyun.credentials.Client newOidcCred() throws Exception {
        // new credential which use rrsa oidc token
        com.aliyun.credentials.models.Config credConf = new com.aliyun.credentials.models.Config();
        credConf.type = "oidc_role_arn";
        credConf.roleArn = System.getenv("ALIBABA_CLOUD_ROLE_ARN");
        credConf.oidcProviderArn = System.getenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN");
        credConf.oidcTokenFilePath = System.getenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE");
        credConf.roleSessionName = "test-rrsa-oidc-token";
        return new com.aliyun.credentials.Client(credConf);
    }
}
