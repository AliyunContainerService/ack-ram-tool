# Permission

Some subcommands require your account to have the required RAM permissions

| Subcommand             | RAM Actions                  |
| ---------------------- | ---------------------------- |
| `rrsa status`          | `cs:DescribeClusterDetail`  |
| `rrsa enable`          | `cs:DescribeClusterDetail` <br/> `cs:ModifyCluster` <br/> `cs:DescribeClusterLogs`  |
| `rrsa disable`          | `cs:DescribeClusterDetail` <br/> `cs:ModifyCluster` <br/> `cs:DescribeClusterLogs` |
| `rrsa associate-role`   | `cs:DescribeClusterDetail` <br/> `ram:GetRole` <br/> `ram:CreateRole` <br/> `ram:UpdateRole` |
