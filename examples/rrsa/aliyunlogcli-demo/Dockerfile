FROM debian:buster as cli
# TARGETPLATFORM

WORKDIR /tmp/
RUN wget https://github.com/AliyunContainerService/ack-ram-tool/releases/download/v0.13.2/ack-ram-tool_0.13.2_Linux_x86_64.tar.gz && \
    tar zxvf ack-ram-tool_0.13.2_Linux_x86_64.tar.gz && chmod +x ack-ram-tool

FROM python:3.11-buster

ENV PIP_INDEX_URL https://mirrors.aliyun.com/pypi/simple/
RUN pip install --upgrade pip && pip install aliyun-log-cli

WORKDIR /usr/bin/
COPY --from=cli /tmp/ack-ram-tool ./ack-ram-tool
