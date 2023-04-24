---
slug: permissions
---

# Permissions

In order to use ack-ram-tool normally, you need to grant the necessary RAM permissions and RBAC permissions for 
the Alibaba Cloud RAM user or RAM role that uses this tool.
The minimum permission information required for each subcommand is shown in the following table:

| Command                            | RAM Permissoins                                                                                                                                                                 | RBAC Permissions |
|------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------|
| `rrsa status`                      | `cs:DescribeClusterDetail`                                                                                                                                                      |                 |
| `rrsa enable`                      | `cs:DescribeClusterDetail` <br/> `cs:ModifyCluster` <br/> `cs:DescribeClusterLogs`                                                                                              |                 |
| `rrsa associate-role`              | `cs:DescribeClusterDetail` <br/> `ram:GetRole` <br/> `ram:CreateRole` <br/> `ram:UpdateRole`                                                                                    |                 |
| `rrsa install-helper-addon`        | `cs:DescribeClusterDetail` <br/> `cs:DescribeClusterAddonsVersion` <br/> `cs:InstallClusterAddons`                                                                              |                 |
| `rrsa assumerole`                  |                                                                                                                                                                                 |                 |
| `rrsa disable`                     | `cs:DescribeClusterDetail` <br/> `cs:ModifyCluster` <br/> `cs:DescribeClusterLogs`                                                                                              |                 |
| `rrsa setup-addon`                 | `cs:DescribeClusterDetail` <br/> `ram:GetRole` <br/> `ram:CreateRole` <br/> `ram:UpdateRole` <br/> `ram:CreatePolicy` <br/> `ram:ListPoliciesForRole` <br/> `ram:AttachPolicyToRole` |                 |
| `rrsa demo`                        |                                                                                                                                                                                 |                 |
| `credential-plugin get-kubeconfig` | `cs:DescribeClusterUserKubeconfig`                                                                                                                                              |                 |
| `credential-plugin get-credential` | `cs:DescribeClusterUserKubeconfig`                                                                                                                                              |                 |
| `credential-plugin get-token`      |                                                                                                                                                                                 |                 |
| `export-credentials`               |                                                                                                                                                                                 |                 |
