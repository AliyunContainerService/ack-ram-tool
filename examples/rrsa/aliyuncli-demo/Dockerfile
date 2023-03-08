FROM registry.cn-hangzhou.aliyuncs.com/acs/ack-ram-tool:1.0.0-dev as cli
# TARGETPLATFORM

WORKDIR /tmp/
RUN wget https://aliyuncli.alicdn.com/aliyun-cli-linux-latest-amd64.tgz && \
    tar zxvf aliyun-cli-linux-latest-amd64.tgz && chmod +x aliyun

FROM alpine:3.17.2

WORKDIR /usr/bin/
COPY --from=cli /usr/bin/ack-ram-tool ./ack-ram-tool
COPY --from=cli /tmp/aliyun ./aliyun

WORKDIR /root/.aliyun/
COPY config.json ./

CMD aliyun cs DescribeClusters
