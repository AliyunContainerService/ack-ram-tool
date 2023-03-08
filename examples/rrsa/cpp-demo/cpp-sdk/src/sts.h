#include <iostream>
#include <cstdlib>
#include <string>
#include <fstream>
#include <json/json.h>

#include <alibabacloud/core/AlibabaCloud.h>
#include <alibabacloud/core/CommonRequest.h>
#include <alibabacloud/core/CommonClient.h>
#include <alibabacloud/core/CommonResponse.h>


using namespace std;

struct Credentials
{
    std::string securityToken;
    std::string accessKeyId;
    std::string accessKeySecret;
    std::string expiration;
    std::string requestId;
};

struct ApiError {
    std::string requestId;
    std::string code;
    std::string message;
};

struct AssumeRoleResult {
    bool success{};
    ApiError error;
    Credentials credentials;
};

Credentials parseCredentials(const std::string &payload)
{
    Json::Reader reader;
    Json::Value value;
    reader.parse(payload, value);

    Credentials cred;
    if(!value["RequestId"].isNull())
        cred.requestId = value["RequestId"].asString();

    auto credentialsNode = value["Credentials"];
    if(!credentialsNode["SecurityToken"].isNull())
        cred.securityToken = credentialsNode["SecurityToken"].asString();
    if(!credentialsNode["Expiration"].isNull())
        cred.expiration = credentialsNode["Expiration"].asString();
    if(!credentialsNode["AccessKeySecret"].isNull())
        cred.accessKeySecret = credentialsNode["AccessKeySecret"].asString();
    if(!credentialsNode["AccessKeyId"].isNull())
        cred.accessKeyId = credentialsNode["AccessKeyId"].asString();

    return cred;
}


std::string readOidcTokenFile(const std::string &path) {
    string line;
    string token;
    ifstream file (path);
    if (file.is_open())
    {
        while ( getline (file,line) )
        {
            if (!line.empty() && line != "\n" && line != "\r\n" && line != "\r") {
                token = line;
                break;
            }
        }
        file.close();
    }
    return token;
}


// TODO: add cache for credentials
AssumeRoleResult AssumeRoleWithOIDC(const std::string &regionId,
                          const std::string &oidcProviderArn,
                          const std::string &oidcTokenFile,
                          const std::string &roleArn,
                          const std::string &sessionName="assume-role-session",
                          const std::string &endpoint="sts.aliyuncs.com",
                          const std::string &durationSeconds="3600") {
    AssumeRoleResult result;
    std::string token = readOidcTokenFile(oidcTokenFile);
    if (token.empty()) {
        std::cout << "read token from file failed: " << oidcTokenFile << std::endl;
        result.success = false;
        result.error.message = "read token from file failed";
        return result;
    }

    AlibabaCloud::ClientConfiguration configuration( regionId );
    AlibabaCloud::Credentials oldCred( "", "");
    AlibabaCloud::CommonClient client( oldCred, configuration );
    AlibabaCloud::CommonRequest request(AlibabaCloud::CommonRequest::RequestPattern::RpcPattern);
    request.setHttpMethod(AlibabaCloud::HttpRequest::Method::Post);
    request.setDomain(endpoint);
    request.setVersion("2015-04-01");
    request.setQueryParameter("Action", "AssumeRoleWithOIDC");
    request.setQueryParameter("RegionId", regionId);
    request.setBodyParameter("OIDCProviderArn", oidcProviderArn);
    request.setBodyParameter("OIDCToken", token);
    request.setBodyParameter("RoleArn", roleArn);
    request.setBodyParameter("DurationSeconds", durationSeconds);
    request.setBodyParameter("RoleSessionName", sessionName);

    auto response = client.commonResponse(request);

    if (response.isSuccess()) {
        string payload = response.result().payload();
        result.credentials = parseCredentials(payload);
        result.success = true;
    } else {
        result.error.requestId = response.error().requestId();
        result.error.message = response.error().errorMessage();
        result.error.code = response.error().errorCode();
    }

    return result;
}
