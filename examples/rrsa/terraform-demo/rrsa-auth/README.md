# Run Terraform with RRSA Auth

```
aliyun/terraform-provider-alicloud >= v1.222.0
```

https://registry.terraform.io/providers/aliyun/alicloud/latest


## Demo


1. Enable RRSA:

```
export CLUSTER_ID=<cluster_id>
ack-ram-tool rrsa enable --cluster-id "${CLUSTER_ID}"
```

2. Install ack-pod-identity-webhook:

```
ack-ram-tool rrsa install-helper-addon --cluster-id "${CLUSTER_ID}"
```

3. Create an RAM Policy:

```
aliyun ram CreatePolicy --PolicyName test-rrsa-create-delete-vpc --PolicyDocument '{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "vpc:DescribeVpcAttribute",
        "vpc:CreateVpc",
        "vpc:DescribeRouteTableList",
        "vpc:DeleteVpc"
      ],
      "Resource": [
        "*"
      ],
      "Condition": {}
    }
  ]
}'
```

4. Associate an RAM Role to the service account and attach the policy to the role:

```
ack-ram-tool rrsa associate-role --cluster-id "${CLUSTER_ID}" \
    --namespace rrsa-demo-terraform \
    --service-account demo-sa \
    --role-name test-rrsa-demo \
    --create-role-if-not-exist \
    --attach-custom-policy test-rrsa-create-delete-vpc
```

5. Deploy demo job:

```
ack-ram-tool credential-plugin get-kubeconfig --cluster-id "${CLUSTER_ID}" > kubeconfig
kubectl --kubeconfig ./kubeconfig apply -f deploy.yaml
```

6. Get logs:

```
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-terraform wait --for=condition=complete job/demo --timeout=240s
kubectl --kubeconfig ./kubeconfig -n rrsa-demo-terraform logs job/demo
```

Outputs:

```
+ terraform init

Initializing the backend...

Initializing provider plugins...
- Reusing previous version of aliyun/alicloud from the dependency lock file
- Using previously-installed aliyun/alicloud v1.222.0

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
+ terraform apply -auto-approve

Terraform used the selected providers to generate the following execution
plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # alicloud_vpc.default will be created
  + resource "alicloud_vpc" "default" {
      + cidr_block            = "10.0.0.0/8"
      + create_time           = (known after apply)
      + description           = "test"
      + id                    = (known after apply)
      + ipv6_cidr_block       = (known after apply)
      + ipv6_cidr_blocks      = (known after apply)
      + name                  = (known after apply)
      + resource_group_id     = (known after apply)
      + route_table_id        = (known after apply)
      + router_id             = (known after apply)
      + router_table_id       = (known after apply)
      + secondary_cidr_blocks = (known after apply)
      + status                = (known after apply)
      + user_cidrs            = (known after apply)
      + vpc_name              = "terraform-with-rrsa-auth-example"
    }

Plan: 1 to add, 0 to change, 0 to destroy.
alicloud_vpc.default: Creating...
alicloud_vpc.default: Creation complete after 6s [id=vpc-XXX]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
+ terraform apply -destroy -auto-approve
alicloud_vpc.default: Refreshing state... [id=vpc-XXX]

Terraform used the selected providers to generate the following execution
plan. Resource actions are indicated with the following symbols:
  - destroy

Terraform will perform the following actions:

  # alicloud_vpc.default will be destroyed
  - resource "alicloud_vpc" "default" {
      - cidr_block            = "10.0.0.0/8" -> null
      - classic_link_enabled  = false -> null
      - create_time           = "2024-04-23T12:07:36Z" -> null
      - description           = "test" -> null
      - enable_ipv6           = false -> null
      - id                    = "vpc-XXX" -> null
      - ipv6_cidr_blocks      = [] -> null
      - name                  = "terraform-with-rrsa-auth-example" -> null
      - resource_group_id     = "rg-XXX" -> null
      - route_table_id        = "vtb-XXX" -> null
      - router_id             = "vrt-XXX" -> null
      - router_table_id       = "vtb-XXX" -> null
      - secondary_cidr_blocks = [] -> null
      - status                = "Available" -> null
      - tags                  = {} -> null
      - user_cidrs            = [] -> null
      - vpc_name              = "terraform-with-rrsa-auth-example" -> null
    }

Plan: 0 to add, 0 to change, 1 to destroy.
alicloud_vpc.default: Destroying... [id=vpc-XXX]
alicloud_vpc.default: Destruction complete after 5s

Apply complete! Resources: 0 added, 0 changed, 1 destroyed.
```
