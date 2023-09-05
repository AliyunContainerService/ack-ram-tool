---
slug: associate-role
sidebar_position: 2
---

# associate-role

Configure RAM roles to allow the use of OIDC tokens representing specific service account identities to assume the RAM roles.


## Usage

```shell
$ ack-ram-tool rrsa associate-role --cluster-id <clusterId> \
  --namespace <namespce> --service-account <serviceAccountName> \
  --role-name <roleName>

? Are you sure you want to associate RAM Role "<roleName>" to service account "<serviceAccountName>" (namespace: "<namespce>")? Yes
2023-04-20T14:30:02+08:00 INFO will change the AssumeRole Policy of RAM Role "<roleName>" with blow content:
{
  "Statement": [
   {
    "Action": "sts:AssumeRole",
    "Condition": {
     "StringEquals": {
      "oidc:aud": "sts.aliyuncs.com",
      "oidc:iss": "https://oidc-ack-***.aliyuncs.com/c132c***",
      "oidc:sub": "system:serviceaccount:<namespce>:<serviceAccountName>"
     }
    },
    "Effect": "Allow",
    "Principal": {
     "Federated": [
      "acs:ram::113***:oidc-provider/ack-rrsa-c132c***"
     ]
    }
   }
  ],
  "Version": "1"
 }

? Are you sure you want to associate RAM Role "test" to service account "sa" (namespace: "test")? Yes
2023-04-20T14:30:04+08:00 INFO Associate RAM Role "test" to service account "sa" (namespace: "test") successfully
```

## Flags

```
Usage:
  ack-ram-tool rrsa associate-role [flags]

Flags:
      --attach-custom-policy string   Attach this custom policy to the RAM Role
      --attach-system-policy string   Attach this system policy to the RAM Role
  -c, --cluster-id string             The cluster id to use
      --create-role-if-not-exist      Create the RAM Role if it does not exist
  -h, --help                          help for associate-role
  -n, --namespace string              The Kubernetes namespace to use
  -r, --role-name string              The RAM Role name to use
  -s, --service-account string        The Kubernetes service account to use

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

| Flag                       | Default | Required | Description                                                         |
|----------------------------|---------|----------|---------------------------------------------------------------------|
| -c, --cluster-id           |         | Yes      | Cluster ID                                                          |
| -n, --namespace            |         | Yes      | namespace，can use `*` to represent all namespaces：`--namespace '*'` |
| -s, --service-account      |         | Yes      | service account                                                     |
| -r, --role-name            |         | Yes      | RAM Role                                                            |
| --create-role-if-not-exist |         |          | auto create an RAM Role if it does not exists                       |
| --attach-system-policy     |         |          | attach a system policy to the role                                  |  
| --attach-custom-policy     |         |          | attach a custom policy to the role                                  |   
