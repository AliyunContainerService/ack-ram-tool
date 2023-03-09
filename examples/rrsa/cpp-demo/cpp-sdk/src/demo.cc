#include <iostream>
#include <cstdlib>
#include <string>
#include <alibabacloud/core/AlibabaCloud.h>
#include <alibabacloud/core/CommonRequest.h>
#include <alibabacloud/core/CommonClient.h>
#include <alibabacloud/core/CommonResponse.h>

#include "sts.h"
#include "cs.h"

using namespace std;


int main() {
    string oidcTokenFile = string(getenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE"));
    string oidcArn = string(getenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN"));
    string roleArn = string(getenv("ALIBABA_CLOUD_ROLE_ARN"));
    AlibabaCloud::InitializeSdk();

    // TODO: add cache
    AssumeRoleResult assumeRoleResult = AssumeRoleWithOIDC("cn-hangzhou", oidcArn, oidcTokenFile, roleArn);
    if (!assumeRoleResult.success) {
        cout << "get credential via assume role failed:" <<endl;
        cout << "RequestId: " << assumeRoleResult.error.requestId << endl;
        cout << "ErrorCode: " << assumeRoleResult.error.code << endl;
        cout << "ErrorMessage: " << assumeRoleResult.error.message << endl;
        return -1;
    }

    // init client
    AlibabaCloud::ClientConfiguration configuration("cn-hangzhou");
    AlibabaCloud::Credentials credential( "", "");
    credential.setAccessKeyId(assumeRoleResult.credentials.accessKeyId);
    credential.setAccessKeySecret(assumeRoleResult.credentials.accessKeySecret);
    credential.setSessionToken(assumeRoleResult.credentials.securityToken);
    AlibabaCloud::CommonClient client( credential, configuration );

    // call open api
    AlibabaCloud::CommonRequest request(AlibabaCloud::CommonRequest::RequestPattern::RoaPattern);
    request.setHttpMethod(AlibabaCloud::HttpRequest::Method::Get);
    request.setDomain("cs.cn-hangzhou.aliyuncs.com");
    request.setVersion("2015-12-15");
    request.setResourcePath("/clusters");
    request.setHeaderParameter("Content-Type", "application/json");

    cout << "call cs.describeClusters via oidc token\n" << endl;

    auto response = client.commonResponse(request);
    if (response.isSuccess()) {
        string payload = response.result().payload();
        DescribeClustersResult result(payload);
        std::vector<DescribeClustersResult::ClusterItem> clusters = result.getClusters();
        for (auto &item : clusters) {
            cout << "name: " << item.name << " ";
            cout << "cluster id: " << item.clusterId << endl;
        }
    } else {
        cout << "error message: " << response.error().errorMessage();
        cout << "error code: " << response.error().errorCode();
        cout << "request id:" << response.error().requestId();
    }

    AlibabaCloud::ShutdownSdk();
    return 0;
}
