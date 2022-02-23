
import com.aliyun.credentials.Client;
import com.aliyun.tea.*;
import com.aliyun.sts20150401.*;
import com.aliyun.sts20150401.models.*;
import com.aliyun.teaopenapi.*;
import com.aliyun.teaopenapi.models.*;
import com.aliyun.credentials.models.Config;

public class Main {
    public static void main(String[] args) throws Exception {
        com.aliyun.credentials.models.Config credConf = new com.aliyun.credentials.models.Config();
        credConf.type = "oidc_role_arn";
        credConf.roleArn = System.getenv("ALIBABA_CLOUD_ROLE_ARN");
        credConf.oidcProviderArn = System.getenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN");
        credConf.oidcTokenFilePath = System.getenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE");
        credConf.roleSessionName = "test-rrsa-oidc-token";
        com.aliyun.credentials.Client cred =  new Client(credConf);


        com.aliyun.teaopenapi.models.Config config = new com.aliyun.teaopenapi.models.Config()
                .setCredential(cred)
                .setEndpoint("sts.aliyuncs.com");

        com.aliyun.sts20150401.Client client = new com.aliyun.sts20150401.Client(config);


        GetCallerIdentityResponse resp = client.getCallerIdentity();
        System.out.println(resp.getBody().toMap().toString());
    }
}
