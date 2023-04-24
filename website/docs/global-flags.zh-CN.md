---
slug: /zh-CN/global-flags
---

# 全局参数

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
