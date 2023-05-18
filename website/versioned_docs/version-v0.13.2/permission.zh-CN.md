---
slug: /zh-CN/permissions
---

# 权限

为了正常使用 ack-ram-tool，您需要为使用改工具的阿里云 RAM 用户或 RAM 角色授予所需的 RAM 权限和 RBAC 权限。
各个子命令所需的最小权限信息如下表所示：

| 子命令                                | RAM 权限                                                                                                                                                                               | RBAC 权限 |
|------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| `rrsa status`                      | `cs:DescribeClusterDetail`                                                                                                                                                           | 无       |
| `rrsa enable`                      | `cs:DescribeClusterDetail` <br/> `cs:ModifyCluster` <br/> `cs:DescribeClusterLogs`                                                                                                   |  无       |
| `rrsa associate-role`              | `cs:DescribeClusterDetail` <br/> `ram:GetRole` <br/> `ram:CreateRole` <br/> `ram:UpdateRole`                                                                                         |  无       |
| `rrsa install-helper-addon`        | `cs:DescribeClusterDetail` <br/> `cs:DescribeClusterAddonsVersion` <br/> `cs:InstallClusterAddons`                                                                                   |  无       |
| `rrsa assumerole`                  | 无                                                                                                                                                                                    |  无       |
| `rrsa disable`                     | `cs:DescribeClusterDetail` <br/> `cs:ModifyCluster` <br/> `cs:DescribeClusterLogs`                                                                                                   |  无       |
| `rrsa setup-addon`                 | `cs:DescribeClusterDetail` <br/> `ram:GetRole` <br/> `ram:CreateRole` <br/> `ram:UpdateRole` <br/> `ram:CreatePolicy` <br/> `ram:ListPoliciesForRole` <br/> `ram:AttachPolicyToRole` |  无       |
| `rrsa demo`                        | 无                                                                                                                                                                                    |  无       |
| `credential-plugin get-kubeconfig` | `cs:DescribeClusterUserKubeconfig`                                                                                                                                                   |  无       |
| `credential-plugin get-credential` | `cs:DescribeClusterUserKubeconfig`                                                                                                                                                   |  无       |
| `credential-plugin get-token`      | 无                                                                                                                                                                                    |  无       |
| `export-credentials`               | 无                                                                                                                                                                                    |  无        |
