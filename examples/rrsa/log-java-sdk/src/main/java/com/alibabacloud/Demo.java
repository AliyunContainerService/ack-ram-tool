package com.alibabacloud;

// com.aliyun:credentials-java >= 0.2.10
import com.aliyun.credentials.Client;

import com.aliyun.openservices.log.*;
import com.aliyun.openservices.log.common.Project;
import com.aliyun.openservices.log.common.auth.*;
import com.aliyun.openservices.log.response.ListProjectResponse;

import java.util.List;

class LogCredentialProvider implements CredentialsProvider {

    private final com.aliyun.credentials.Client cred;

    public LogCredentialProvider(com.aliyun.credentials.Client cred) {
        this.cred = cred;
    }

    @Override
    public Credentials getCredentials() {
        String ak = cred.getAccessKeyId();
        String sk = cred.getAccessKeySecret();
        String token = cred.getSecurityToken();
        return new DefaultCredentials(ak, sk, token);
    }
}

class TestLogSDK {

    public void CallAPI(com.aliyun.credentials.Client cred) throws Exception {
        // new provider
        LogCredentialProvider provider = new LogCredentialProvider(cred);
        String endpoint = "cn-hangzhou.log.aliyuncs.com";
        // init client
        com.aliyun.openservices.log.Client client = new com.aliyun.openservices.log.Client(endpoint, provider);

        // call api
        ListProjectResponse response = client.ListProject();
        System.out.println("call Log.ListProject via oidc token success:\n");
        for (Project project : response.getProjects()) {
            System.out.println(" - " + project.getProjectName());
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

        // test Log sdk (https://github.com/aliyun/aliyun-Log-java-sdk) use rrsa oidc token
        System.out.println("test Log sdk use rrsa oidc token");
        TestLogSDK logsdk = new TestLogSDK();
        logsdk.CallAPI(cred);
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
