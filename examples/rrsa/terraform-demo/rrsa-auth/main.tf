terraform {
  required_providers {
    alicloud = {
      source = "aliyun/alicloud"
      version = ">=1.220.0"
    }
  }
}

provider "alicloud" {
  // https://registry.terraform.io/providers/aliyun/alicloud/latest/docs#assuming-a-ram-role-with-oidc
  assume_role_with_oidc {
    role_session_name = "terraform-with-rrsa-auth-example"
  }

  region = "cn-beijing"
}

resource "alicloud_vpc" "default" {
  description = "test"
  cidr_block  = "10.0.0.0/8"
  vpc_name    = "terraform-with-rrsa-auth-example"
}
