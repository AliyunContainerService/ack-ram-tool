package env

import "os"

var (
	accessKeyIdEnvs = []string{
		envNewSdkAccessKeyId,
		envOldSdkAccessKeyID,
		envAccAlibabaCloudAccessKeyId,
	}

	accessKeySecretEnvs = []string{
		envNewSdkAccessKeySecret,
		envOldSdkAccessKeySecret,
		envAccAlibabaCloudAccessKeySecret,
	}

	securityTokenEnvs = []string{
		envNewSdkSecurityToken,
		envOldSdkAccessKeyStsToken,
		envAccAlibabaCloudSecurityToken,
	}

	credentialsURIEnvs = []string{
		envNewSdkCredentialsURI,
	}
)

func GetAccessKeyId() string {
	for _, key := range accessKeyIdEnvs {
		v := os.Getenv(key)
		if v != "" {
			return v
		}
	}
	return ""
}

func GetAccessKeySecret() string {
	for _, key := range accessKeySecretEnvs {
		v := os.Getenv(key)
		if v != "" {
			return v
		}
	}
	return ""
}

func GetSecurityToken() string {
	for _, key := range securityTokenEnvs {
		v := os.Getenv(key)
		if v != "" {
			return v
		}
	}
	return ""
}

func GetCredentialsURI() string {
	for _, key := range credentialsURIEnvs {
		v := os.Getenv(key)
		if v != "" {
			return v
		}
	}
	return ""
}
