package env

// https://github.com/aliyun/credentials-go
const (
	envNewSdkAccessKeyId     = "ALIBABA_CLOUD_ACCESS_KEY_ID"
	envNewSdkAccessKeySecret = "ALIBABA_CLOUD_ACCESS_KEY_SECRET" // #nosec G101
	envNewSdkSecurityToken   = "ALIBABA_CLOUD_SECURITY_TOKEN"    // #nosec G101
	envNewSdkRoleSessionName = "ALIBABA_CLOUD_ROLE_SESSION_NAME"

	envNewSdkCredentialsURI = "ALIBABA_CLOUD_CREDENTIALS_URI" // #nosec G101

	envNewSdkCredentialFile = "ALIBABA_CLOUD_CREDENTIALS_FILE" // #nosec G101

	EnvRoleSessionName = envNewSdkRoleSessionName
)
