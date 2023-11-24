# RBAC


## scan-user-permissions

扫描指定集群中存在的 RAM 用户和角色的 RBAC bindings.

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

## cleanup-user-permissions

清理指定集群中指定的 RAM 用户和角色的 RBAC bindings.

```shell

# <用户ID> 可以从上面的 scan 命令的结果中获取
$ ack-ram-tool rbac cleanup-user-permissions -c <集群ID>  -u <用户ID>
Start to scan users and bindings
Will cleanup RBAC bindings as blow:
UID                                UserType  UserName  Binding                                                         
300******************** (deleted)  RamRole             RoleBinding/kube-system/300********************-heapster-rolebinding  
300******************** (deleted)  RamRole             RoleBinding/arms-prom/300********************-arms-prom-rolebinding   
300******************** (deleted)  RamRole             RoleBinding/default/300********************-default-rolebinding       
300******************** (deleted)  RamRole             ClusterRoleBinding/-/300********************-clusterrolebinding       
? Are you sure you want to cleanup these bindings? Yes
start to cleanup binding: RoleBinding/kube-system/300********************-heapster-rolebinding
finished cleanup binding: RoleBinding/kube-system/300********************-heapster-rolebinding
start to cleanup binding: RoleBinding/arms-prom/300********************-arms-prom-rolebinding
finished cleanup binding: RoleBinding/arms-prom/300********************-arms-prom-rolebinding
start to cleanup binding: RoleBinding/default/300********************-default-rolebinding
finished cleanup binding: RoleBinding/default/300********************-default-rolebinding
start to cleanup binding: ClusterRoleBinding/-/300********************-clusterrolebinding
finished cleanup binding: ClusterRoleBinding/-/300********************-clusterrolebinding
all bindings have been cleanup

```
