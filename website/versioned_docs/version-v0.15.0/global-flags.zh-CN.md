---
slug: /zh-CN/global-flags
---

# 全局参数

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

| 参数名称                            | 默认值                                                     | 说明                                    |
|---------------------------------|---------------------------------------------------------|---------------------------------------|
| -y, --assume-yes                | false                                                   | 为 true 时，程序将自动执行，不再询问是否继续操作           |
| --log-level                     | info                                                    | 日志级别                                  |
| --ignore-aliyun-cli-credentials | false                                                   | 为 true 时，查找凭证信息时将忽略 aliyun cli 的配置文件  |
| --ignore-env-credentials        | false                                                   | 为 true 时，查找凭证信息时将忽略环境变量中的凭证信息         |
| --profile-file                  | `~/.aliyun/config.json` 或 `~/.alibabacloud/credentials` | 指定凭证配置文件                              |
| --profile-name                  |                                                        | 当使用 aliyun cli 配置文件时，使用指定的配置集中定义的凭证配置 |
| --region-id                     | cn-hangzhou                                             | 访问 open api 时使用的 region 信息            |
| -v, --verbose                     | false                                                   | 快速启用 debug 模式                         |
