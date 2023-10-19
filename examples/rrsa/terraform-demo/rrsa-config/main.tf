provider "alicloud" {
}

variable "k8s_name_prefix" {
  description = "The name prefix used to create ASK cluster."
  default     = "ask-rrsa-example"
}

resource "random_uuid" "this" {}


locals {
  k8s_name_ask = substr(join("-", [var.k8s_name_prefix,"ask"]), 0, 63)
  new_vpc_name = "tf-vpc-172-16"
  new_vsw_name = "tf-vswitch-172-16-0"
}

data "alicloud_zones" "default" {
  available_resource_creation = "VSwitch"
}

resource "alicloud_vpc" "vpc" {
  vpc_name   = local.new_vpc_name
  cidr_block = "172.16.0.0/12"
}

resource "alicloud_vswitch" "vsw" {
  vswitch_name      = local.new_vsw_name
  vpc_id          = alicloud_vpc.vpc.id
  cidr_block      = cidrsubnet(alicloud_vpc.vpc.cidr_block, 8, 8)
  zone_id      = data.alicloud_zones.default.zones[0].id
}


resource "alicloud_cs_serverless_kubernetes" "serverless" {
  name                           = local.k8s_name_ask
  version                        = "1.26.3-aliyun.1"
  cluster_spec                   = "ack.pro.small"
  vpc_id                         = alicloud_vpc.vpc.id
  vswitch_ids                    = split(",", join(",", alicloud_vswitch.vsw.*.id))
  new_nat_gateway                = false
  endpoint_public_access_enabled = false
  deletion_protection            = false
  load_balancer_spec             = "slb.s2.small"
  time_zone                      = "Asia/Shanghai"
  service_cidr                   = "10.13.0.0/16"
  service_discovery_types        = ["CoreDNS"]

  # Enable RRSA
  enable_rrsa                    = true
}


# k8s service account info
variable "k8s_namespace" {
    default       = "test-rrsa-ns"
}
variable "k8s_service_account" {
    default = "foo-bar-manager-sa"
}

# Create a new RAM Role.
resource "alicloud_ram_role" "role" {
  name        = "rrsa-demo-${alicloud_cs_serverless_kubernetes.serverless.id}"
  document    = <<EOF
  {
    "Statement": [
      {
        "Action": "sts:AssumeRole",
        "Condition": {
          "StringEquals": {
            "oidc:aud": "sts.aliyuncs.com",
            "oidc:iss": "${alicloud_cs_serverless_kubernetes.serverless.rrsa_metadata[0].rrsa_oidc_issuer_url}",
            "oidc:sub": "system:serviceaccount:${var.k8s_namespace}:${var.k8s_service_account}"
          }
        },
        "Effect": "Allow",
        "Principal": {
          "Federated": [
            "${alicloud_cs_serverless_kubernetes.serverless.rrsa_metadata[0].ram_oidc_provider_arn}"
          ]
        }
      }
    ],
    "Version": "1"
  }
  EOF
  description = "this is a role for rrsa demo."
  force       = true
}

# Create a new RAM Policy.
resource "alicloud_ram_policy" "policy" {
  policy_name        = "rrsa-demo-policy-demo"
  policy_document    = <<EOF
  {
    "Statement": [
      {
        "Action": [
          "oss:GetObject"
        ],
        "Effect": "Allow",
        "Resource": [
          "acs:oss:*:*:my-foo-bar-bucket/*"
        ]
      }
    ],
      "Version": "1"
  }
  EOF
  description = "this is a policy test"
  force       = true
}

# Attach Policy to the Role.
resource "alicloud_ram_role_policy_attachment" "attach" {
  policy_name = alicloud_ram_policy.policy.policy_name
  policy_type = alicloud_ram_policy.policy.type
  role_name   = alicloud_ram_role.role.name
}

