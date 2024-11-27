package exportcredentials

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
)

var (
	accessKeyIdEnvs = []string{
		"ALIBABA_CLOUD_ACCESS_KEY_ID",
		"ALICLOUD_ACCESS_KEY",
		"ALIBABACLOUD_ACCESS_KEY_ID",
		"ALIYUN_LOG_CLI_ACCESSID",
	}

	accessKeySecretEnvs = []string{
		"ALIBABA_CLOUD_ACCESS_KEY_SECRET",
		"ALICLOUD_SECRET_KEY",
		"ALIBABACLOUD_ACCESS_KEY_SECRET",
		"ALIYUN_LOG_CLI_ACCESSKEY",
	}
	stsTokenEnvs = []string{
		"ALIBABA_CLOUD_SECURITY_TOKEN",
		"ALICLOUD_SECURITY_TOKEN",
		"ALIBABACLOUD_SECURITY_TOKEN",
		"ALICLOUD_ACCESS_KEY_STS_TOKEN",
		"ALIYUN_LOG_CLI_STS_TOKEN",
	}
)

func runUserCommands(ctx context.Context, cred Credentials, args []string,
	stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	cmd := exec.CommandContext(ctx, args[0], args[1:]...) // #nosec G204
	if stdin == nil {
		stdin = os.Stdin
	}
	if stdout == nil {
		stdout = os.Stdout
	}
	if stderr == nil {
		stderr = os.Stderr
	}
	envs := getCredentialsEnvsWithCurrentEnvs(cred)
	envs = append(envs, "ALIBABACLOUD_IGNORE_PROFILE=TRUE")
	cmd.Env = envs
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return cmd.Run()
}

func getCredentialsEnvsWithCurrentEnvs(cred Credentials) []string {
	var envs []string
	envs = append(envs, os.Environ()...)
	envs = append(envs, getCredentialsEnvs(cred)...)
	return envs
}

func getCredentialsEnvs(cred Credentials) []string {
	var envs []string
	for _, key := range accessKeyIdEnvs {
		envs = append(envs, fmt.Sprintf("%s=%s", key, cred.AccessKeyId))
	}
	for _, key := range accessKeySecretEnvs {
		envs = append(envs, fmt.Sprintf("%s=%s", key, cred.AccessKeySecret))
	}
	if cred.SecurityToken != "" {
		for _, key := range stsTokenEnvs {
			envs = append(envs, fmt.Sprintf("%s=%s", key, cred.SecurityToken))
		}
	}

	return envs
}
