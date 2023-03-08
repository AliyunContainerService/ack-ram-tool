
#include <aws/core/Aws.h>
#include <aws/core/auth/AWSCredentialsProvider.h>
#include <aws/core/auth/STSCredentialsProvider.h>
#include <aws/core/utils/logging/LogLevel.h>
#include <aws/s3/S3Client.h>
#include <aws/s3/model/ListObjectsRequest.h>
#include <iostream>
#include "AliyunCredentialsProvider.h"

using namespace Aws;

int main(int argc, char *argv[])
{
    if (argc < 2 || argc > 3) {
        std::cerr << "Usage: " << argv[0] << "[<region>] <bucket-name>" << std::endl;
        return 1;
    }
    std::string region, bucket;
    if (argc == 2) {
        bucket = argv[1];
        region = "cn-hangzhou";
    } else {
        region = argv[1];
        bucket = argv[2];
    }

    //The Aws::SDKOptions struct contains SDK configuration options.
    //An instance of Aws::SDKOptions is passed to the Aws::InitAPI and 
    //Aws::ShutdownAPI methods.  The same instance should be sent to both methods.
    SDKOptions options;
    options.loggingOptions.logLevel = Utils::Logging::LogLevel::Debug;
    
    //The AWS SDK for C++ must be initialized by calling Aws::InitAPI.
    InitAPI(options); 
    {
        Aws::Client::ClientConfiguration awsCC;
        awsCC.region = "oss-"+region;
        awsCC.endpointOverride = "oss-" + region + ".aliyuncs.com";

        auto aliyunCredentialsProvider = Aws::MakeShared<Aws::Auth::AliyunSTSAssumeRoleWebIdentityCredentialsProvider>("AliyunSTSAssumeRoleWebIdentityCredentialsProvider");
        std::shared_ptr<S3::S3EndpointProviderBase> endpointProvider = Aws::MakeShared<S3::S3EndpointProvider>("S3EndpointProvider");
        S3::S3Client client = S3::S3Client(aliyunCredentialsProvider, endpointProvider, awsCC);

        auto request = S3::Model::ListObjectsRequest();
        request.WithBucket(bucket);

        auto outcome = client.ListObjects(request);

        if (!outcome.IsSuccess()) {
            std::cerr << "Error: ListObjects: " <<
                    outcome.GetError().GetMessage() << std::endl;
        }
        else {
            Aws::Vector<Aws::S3::Model::Object> objects =
                    outcome.GetResult().GetContents();

            for (Aws::S3::Model::Object &object: objects) {
                std::cout << object.GetKey() << std::endl;
            }
        }
    }

    //Before the application terminates, the SDK must be shut down. 
    ShutdownAPI(options);
    return 0;
}