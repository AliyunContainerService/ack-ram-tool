---
slug: get-token
sidebar_position: 3
---

# get-token

Integrate ack-ram-authenticator to obtain the [ExecCredential](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#client-go-credential-plugins) token used to access the API server.

## Usage

```shell
$ ack-ram-tool credential-plugin get-token --cluster-id <clusterId>

{
 "kind": "ExecCredential",
 "apiVersion": "client.authentication.k8s.io/v1beta1",
 "spec": {
  "interactive": false
 },
 "status": {
  "token": "k8s-ack-v1.aHR0cHM6Ly9zd***"
 }
}
```

## Flags

```
Usage:
  ack-ram-tool credential-plugin get-token [flags]

Flags:
      --api-version string   v1 or v1beta1 (default "v1beta1")
  -c, --cluster-id string    The cluster id to use
  -h, --help                 help for get-token
      --role-arn string      Assume an RAM Role ARN when send request or sign token

Global Flags:
  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively
      --ignore-aliyun-cli-credentials   don't try to parse credentials from config.json of aliyun cli
      --ignore-env-credentials          don't try to parse credentials from environment variables
      --log-level string                log level: info, debug, error (default "info")
      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)
      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli
      --region-id string                The region to use (default "cn-hangzhou")
```

Descriptions：

| Flag             | Default | Required | Description                                                                                                                                                                          |
|------------------|---------|----------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| -c, --cluster-id |         | Yes      | Cluster ID                                                                                                                                                                           |
| --api-version    | v1beta1 |          | Specify which version of apiVersion to use in the returned data. `v1beta1` represents `client.authentication.k8s.io/v1beta1`, and `v1` represents `client.authentication.k8s.io/v1`. |
| --role-arn       |         |          | Assume an RAM Role ARN when send request or sign token                                                                                                                               |
