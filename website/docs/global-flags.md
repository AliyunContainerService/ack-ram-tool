---
slug: global-flags
---

# Global Flags

```
Global Flags:
  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively
      --ignore-aliyun-cli-credentials   don't try to parse credentials from config.json of aliyun cli
      --ignore-env-credentials          don't try to parse credentials from environment variables
      --log-level string                log level: info, debug, error (default "info")
      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials)
      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli
      --region-id string                The region to use (default "cn-hangzhou")
      -v, --verbose                     Make the operation more talkative
```

## 参数说明

| Flag                            | Default                                                 | Description                                                                                                                         |
|---------------------------------|---------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------|
| -y, --assume-yes                | false                                                   | When set to true, the program will automatically execute without asking whether to continue the operation                           |
| --log-level                     | info                                                    | Log level                                                                                                                           |
| --ignore-aliyun-cli-credentials | false                                                   | When set to true, the aliyun cli configuration file will be ignored when searching for credential information                       |
| --ignore-env-credentials        | false                                                   | When set to true, the credential information in the environment variables will be ignored when searching for credential information |
| --profile-file                  | `~/.aliyun/config.json` 或 `~/.alibabacloud/credentials` | Specify the credential configuration file                                                                                           |
| --profile-name                  | no default                                              | When using the aliyun cli configuration file, use the credential configuration defined in the specified configuration set           |
| --region-id                     | cn-hangzhou                                             | Region information used when accessing OpenAPI                                                                                      |
| -v, --verbose                   | false                                                   | Quickly enable debug mode                                                                                                           |
