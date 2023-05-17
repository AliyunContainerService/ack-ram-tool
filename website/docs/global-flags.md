---
slug: global-flags
---

# Global Flags

```
Global Flags:
  -y, --assume-yes                      Automatic yes to prompts; assume "yes" as answer to all prompts and run non-interactively (env: "ACK_RAM_TOOL_ASSUME_YES")
      --ignore-aliyun-cli-credentials   don't try to parse credentials from config.json of aliyun cli (env: "ACK_RAM_TOOL_IGNORE_ALIYUN_CLI_CREDENTIALS")
      --ignore-env-credentials          don't try to parse credentials from environment variables (env: "ACK_RAM_TOOL_IGNORE_ENV_CREDENTIALS")
      --log-level string                log level: info, debug, error (default "info") (env: "ACK_RAM_TOOL_LOG_LEVEL")
      --profile-file string             Path to credential file (default: ~/.aliyun/config.json or ~/.alibabacloud/credentials) (env: "ACK_RAM_TOOL_PROFILE_FILE")
      --profile-name string             using this named profile when parse credentials from config.json of aliyun cli (env: "ACK_RAM_TOOL_PROFIL_ENAME")
      --region-id string                The region to use (default "cn-hangzhou") (env: "ACK_RAM_TOOL_REGION_ID")
  -v, --verbose                         Make the operation more talkative
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
