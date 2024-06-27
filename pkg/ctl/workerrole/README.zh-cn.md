# worker-role

## 配置命令行使用的凭证

可以参考如下文档：[凭证信息](https://aliyuncontainerservice.github.io/ack-ram-tool/zh-CN/getting-started#%E5%87%AD%E8%AF%81%E4%BF%A1%E6%81%AF)


## 执行命令行时使用的账号凭证所需的权限

执行命令行时使用的账号凭证需要至少被授予如下 **RAM 权限和 RBAC 权限** 。


### RAM 权限

至少需要如下 RAM 权限策略：

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "cs:DescribeClusterAddonsVersion",
        "cs:DescribeClusterUserKubeconfig"
      ],
      "Resource": "*"
    }
  ]
}
```

### RBAC 权限

至少需要如下 RBAC 权限：

* 目标集群的管理员权限。


## scan-addon

辅助列出目标集群中已安装的系统组件所依赖的 worker role 被授予的 ram 权限列表：

```shell
$ ack-ram-tool worker-role scan-addon -c <集群ID>
2024-06-27T16:22:18+08:00 INFO start to scan cluster cd0daxxx
2024-06-27T16:22:38+08:00 WARN [terway-eniip.kube-system.DaemonSet] terway should be updated, which depends on [AliyunCSManagedNetworkRole]
2024-06-27T16:22:38+08:00 WARN [aliyun-acr-credential-helper.kube-system.Deployment] aliyun-acr-credential-helper should be updated, which depends on [AliyunCSManagedAcrRole]
2024-06-27T16:22:38+08:00 INFO Summary of the Scan Result:
● Addon terway needs to be updated, which depends on AliyunCSManagedNetworkRolePolicy
● Addon aliyun-acr-credential-helper needs to be updated, which depends on AliyunCSManagedAcrRolePolicy
● These addons need RAM actions as follows:

  "cr:GetAuthorizationToken",
  "cr:ListInstanceEndpoint",
  "cr:PullRepository",
  "ecs:AssignIpv6Addresses",
  "ecs:AssignPrivateIpAddresses",
  "ecs:AttachNetworkInterface",
  "ecs:CreateNetworkInterface",
  "ecs:DeleteNetworkInterface",
  "ecs:DescribeInstanceAttribute",
  "ecs:DescribeInstanceTypes",
  "ecs:DescribeInstances",
  "ecs:DescribeNetworkInterfaces",
  "ecs:DetachNetworkInterface",
  "ecs:ModifyNetworkInterfaceAttribute",
  "ecs:UnassignIpv6Addresses",
  "ecs:UnassignPrivateIpAddresses",
  "vpc:DescribeVSwitches"

```
输出结果中的【These addons need RAM actions as follows】
之后的内容是这些系统组件所依赖的 RAM 权限 Action 列表（每个集群的输出不一定相同）。

复制输出的 Action 列表或者通过下面的方法将列出的 RAM Action 保存到文件中：

```shell
$ ack-ram-tool worker-role scan-addon -c <集群ID> | tee actions.txt
```
