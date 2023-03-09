FROM registry.cn-hangzhou.aliyuncs.com/acs/ack-ram-tool:0.13.0-dev as cli
# TARGETPLATFORM

FROM python:3.11-buster

ENV PIP_INDEX_URL https://mirrors.aliyun.com/pypi/simple/
RUN pip install --upgrade pip && pip install aliyun-log-cli

WORKDIR /usr/bin/
COPY --from=cli /usr/bin/ack-ram-tool ./ack-ram-tool
