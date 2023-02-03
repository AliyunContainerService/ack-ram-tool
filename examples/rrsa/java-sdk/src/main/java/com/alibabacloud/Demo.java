package com.alibabacloud;

// com.aliyun:credentials-java >= 0.2.10
import com.aliyun.credentials.Client;
import com.aliyun.cs20151215.models.DescribeClustersRequest;

// only for oss sdk
import com.aliyun.cs20151215.models.DescribeClustersResponse;
import com.aliyun.oss.ClientBuilderConfiguration;
import com.aliyun.oss.OSS;
import com.aliyun.oss.common.auth.*;
import com.aliyun.oss.OSSClientBuilder;
import com.aliyun.oss.model.Bucket;

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

class OSSCredentialProvider implements CredentialsProvider {

    private final com.aliyun.credentials.Client cred;

    public OSSCredentialProvider(com.aliyun.credentials.Client cred) {
        this.cred = cred;
    }

    public void setCredentials(Credentials creds) {
    }

    @Override
    public Credentials getCredentials() {
        String ak = cred.getAccessKeyId();
        String sk = cred.getAccessKeySecret();
        String token = cred.getSecurityToken();
        return new DefaultCredentials(ak, sk, token);
    }
}

class TestOSSSDK {

    public void CallAPI(com.aliyun.credentials.Client cred) throws Exception {
        // new provider
        OSSCredentialProvider provider = new OSSCredentialProvider(cred);
        String endpoint = "https://oss-cn-hangzhou.aliyuncs.com";
        // new client config
        ClientBuilderConfiguration conf = new ClientBuilderConfiguration();

        // init client
        OSS ossClient = new OSSClientBuilder().build(endpoint, provider, conf);

        // call api
        List<Bucket> buckets = ossClient.listBuckets();
        System.out.println("call oss.listBuckets via oidc token success:\n");
        for (Bucket bucket : buckets) {
            System.out.println(" - " + bucket.getName());
        }
        System.out.println();

        ossClient.shutdown();
    }

}


public class Demo {

    public static void main(String[] args) throws Exception {
        // 两种方式都可以
        com.aliyun.credentials.Client cred = new Client();
        // or
        // com.aliyun.credentials.Client cred = newOidcCred();


        // test open api sdk (https://github.com/aliyun/alibabacloud-java-sdk) use rrsa oidc token
        System.out.println("\n");
        System.out.println("test open api sdk use rrsa oidc token");
        TestOpenAPISDK openapiSdk = new TestOpenAPISDK();
        openapiSdk.CallAPI(cred);

        // test oss sdk (https://github.com/aliyun/aliyun-oss-java-sdk) use rrsa oidc token
        if (System.getenv("TEST_OSS_SDK") != null && System.getenv("TEST_OSS_SDK").equals("true")) {
            System.out.println("\n");
            System.out.println("test oss sdk use rrsa oidc token");
            TestOSSSDK osssdk = new TestOSSSDK();
            osssdk.CallAPI(cred);
        }
    }

    static com.aliyun.credentials.Client newOidcCred() throws Exception {
        // new credential which use rrsa oidc token
        com.aliyun.credentials.models.Config credConf = new com.aliyun.credentials.models.Config();
        credConf.type = "oidc_role_arn";
        credConf.roleArn = System.getenv("ALIBABA_CLOUD_ROLE_ARN");
        credConf.oidcProviderArn = System.getenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN");
        credConf.oidcTokenFilePath = System.getenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE");
        credConf.roleSessionName = "test-rrsa-oidc-token";
        return new Client(credConf);
    }
}
