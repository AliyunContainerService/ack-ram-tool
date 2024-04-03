FROM debian:11-slim as builder

WORKDIR /app
RUN apt update && apt install wget unzip -y && \
    wget https://releases.hashicorp.com/terraform/1.7.5/terraform_1.7.5_linux_amd64.zip && \
    unzip terraform_1.7.5_linux_amd64.zip

FROM alpine:3.19.1

COPY --from=builder /app/terraform /usr/bin/terraform

WORKDIR /app

COPY main.tf main.tf
COPY run.sh run.sh

RUN terraform init

CMD ["/bin/sh", "/app/run.sh"]
