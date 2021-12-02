# ack-ram-tool

A command line utility and library for using RAM in ACK.


## Credential

You can use ~/.alibabacloud/credentials file:

```
$ cat ~/.alibabacloud/credentials

[default]
type = access_key
access_key_id = foo
access_key_secret = bar
```

Or environment variables:

```
$ export ALIBABA_CLOUD_ACCESS_KEY_Id=foo
$ export ALIBABA_CLOUD_ACCESS_KEY_SECRET=bar
```

## Usage


### RAM Roles for Service Accounts(RRSA)


```
$ ack-ram-tool rrsa status -c <clusterId>
```

```
$ ack-ram-tool rrsa enable -c <clusterId>
```

```
$ ack-ram-tool rrsa disable -c <clusterId>
```

```
$ ack-ram-tool rrsa associate-role -c <clusterId> -r <roleName> \
    -n <namespace> -s <serviceAccount>
```

```
$ ack-ram-tool rrsa assume-role -r <roleArn> -p <oidcProviderArn> \
    -t <oidcTokenFile>
```
