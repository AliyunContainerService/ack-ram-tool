#!/usr/bin/env sh

set -xe

terraform init
terraform apply -auto-approve
terraform apply -destroy -auto-approve
