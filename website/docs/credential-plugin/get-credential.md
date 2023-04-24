---
slug: get-credential
sidebar_position: 2
---

# get-credential

Get the ExecCredential certificate data used to access the API server.

It has the following features:

* Automatically obtains a new certificate before the certificate expires
* Supports using temporary certificate


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
  "expirationTimestamp": "2023-04-20T09:29:06Z",
  "clientCertificateData": "-----BEGIN CERTIFICATE-----\nMIID***\n-----END CERTIFICATE-----\n",
  "clientKeyData": "-----BEGIN RSA PRIVATE KEY-----\nMIIE***\n-----END RSA PRIVATE KEY-----\n"
 }
}
```

## Flags

```
Usage:
  ack-ram-tool credential-plugin get-credential [flags]

Flags:
      --api-version string            v1 or v1beta1 (default "v1beta1")
  -c, --cluster-id string             The cluster id to use
      --credential-cache-dir string   Directory to cache credential (default "~/.kube/cache/ack-ram-tool/credential-plugin")
      --expiration duration           The credential expiration (default 3h0m0s)
  -h, --help                          help for get-credential

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

| Flag                   | Default                                        | Required | Description                                                                                                                                                                                               |
|------------------------|------------------------------------------------|----------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| -c, --cluster-id       |                                                | Yes      | Cluster ID                                                                                                                                                                                                |
| --api-version          | v1beta1                                        |          | Specify which version of apiVersion to use in the returned data. `v1beta1` represents `client.authentication.k8s.io/v1beta1`, and `v1` represents `client.authentication.k8s.io/v1`.                      |
| --expiration           | 3h0m0s                                         |          | Specify the certificate expiration time. When it is 0, it means not to use a temporary certificate but to use a longer valid certificate (the expiration time is automatically determined by the server). |
| --credential-cache-dir | `~/.kube/cache/ack-ram-tool/credential-plugin` |          | Directory used to cache the certificate                                                                                                                                                                   |
