# RBAC

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
        "cs:DescribeClusterUserKubeconfig",
        "cs:DescribeClustersV1",
        "cs:GetClusters",
        "cs:DescribeClusterDetail"
      ],
      "Resource": "*"
    },
    {
      "Effect": "Allow",
      "Action": [
        "ram:ListUsers",
        "ram:ListRoles"
      ],
      "Resource": "*"
    }
  ]
}
```

### RBAC 权限

至少需要如下 RBAC 权限：

* 目标集群的管理员权限。


## scan-user-permissions

* 扫描指定集群中存在的 RAM 用户和角色的 RBAC bindings.

```shell

# 默认只输出已删除 RAM 用户和角色
$ ack-ram-tool rbac scan-user-permissions -c <集群ID>

UID                     UserType  UserName  Binding                                                    
2432******** (deleted)  RamUser             ClusterRoleBinding/-/24*****-clusterrolebinding 


# 可以通过 -A 显示所有用户和角色
$ ack-ram-tool rbac scan-user-permissions -A -c <集群ID>

UID                     UserType  UserName  Binding                                                    
2432******** (deleted)  RamUser             ClusterRoleBinding/-/24*****-clusterrolebinding 
2342********            RamUser   foobar    ClusterRoleBinding/-/23*****-clusterrolebinding 

```

* 扫描所有集群中存在的 RAM 用户和角色的 RBAC bindings.

```shell

# 默认只输出已删除 RAM 用户和角色
$ ack-ram-tool rbac scan-user-permissions -c all

ClusterId: cbbXXX
UID                     UserType  UserName  Binding                                                    
2432******** (deleted)  RamUser             ClusterRoleBinding/-/24*****-clusterrolebinding 


# 可以通过 -A 显示所有用户和角色
$ ack-ram-tool rbac scan-user-permissions -A -c all

ClusterId: cbbXXX
UID                     UserType  UserName  Binding                                                    
2432******** (deleted)  RamUser             ClusterRoleBinding/-/24*****-clusterrolebinding 
2342********            RamUser   foobar    ClusterRoleBinding/-/23*****-clusterrolebinding 

```

## cleanup-user-permissions

* 清理指定集群中指定的 RAM 用户和角色的 RBAC bindings:

```shell

# <用户ID> 可以从上面的 scan 命令的结果中获取
$ ack-ram-tool rbac cleanup-user-permissions -c <集群ID>  -u <用户ID>

Start to scan users and bindings
Will cleanup RBAC bindings as below:
UID               UserType  UserName  Binding                                                         
300*** (deleted)  RamRole             RoleBinding/kube-system/300****-heapster-rolebinding  
300*** (deleted)  RamRole             RoleBinding/arms-prom/300****-arms-prom-rolebinding   
300*** (deleted)  RamRole             RoleBinding/default/300****-default-rolebinding       
300*** (deleted)  RamRole             ClusterRoleBinding/-/300***-clusterrolebinding       
? Are you sure you want to cleanup these bindings? Yes
start to cleanup binding: RoleBinding/kube-system/300***-heapster-rolebinding
finished cleanup binding: RoleBinding/kube-system/300***-heapster-rolebinding
start to cleanup binding: RoleBinding/arms-prom/300***-arms-prom-rolebinding
finished cleanup binding: RoleBinding/arms-prom/300***-arms-prom-rolebinding
start to cleanup binding: RoleBinding/default/300***-default-rolebinding
finished cleanup binding: RoleBinding/default/300***-default-rolebinding
start to cleanup binding: ClusterRoleBinding/-/300***-clusterrolebinding
finished cleanup binding: ClusterRoleBinding/-/300***-clusterrolebinding
all bindings have been cleanup

```


* 清理指定集群中所有已删除的 RAM 用户和角色的 RBAC bindings:

```shell

$ ack-ram-tool rbac cleanup-user-permissions -c <集群ID> --all-deleted-users

Start to scan users and bindings
Will cleanup RBAC bindings as below:
UID               UserType  UserName  Binding                                                         
300*** (deleted)  RamRole             ClusterRoleBinding/-/300***-clusterrolebinding       
? Are you sure you want to cleanup these bindings? Yes
start to backup binding: ClusterRoleBinding/-/300***-clusterrolebinding
the origin binding ClusterRoleBinding/-/300***-clusterrolebinding have been backed up to file cbbXXX/ClusterRoleBinding--300***-clusterrolebinding.json
start to delete binding: ClusterRoleBinding/-/300***-clusterrolebinding
deleted binding: ClusterRoleBinding/-/300***-clusterrolebinding

all bindings have been cleanup
