FROM debian:11-slim as builder

WORKDIR /app
RUN wget https://releases.hashicorp.com/terraform/1.7.5/terraform_1.7.5_linux_amd64.zip && \
    tar zxvf terraform_1.7.5_linux_amd64.zip

FROM alpine:3.19.1

COPY --from=builder /app/terraform /usr/bin/terraform

WORKDIR /app

COPY main.tf main.tf
COPY run.sh run.sh

RUN terraform init

COMMAND ["/bin/sh", "/app/run.sh"]
