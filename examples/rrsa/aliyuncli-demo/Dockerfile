FROM alpine:3.17.2 as cli
# TARGETPLATFORM

WORKDIR /tmp/
RUN wget https://aliyuncli.alicdn.com/aliyun-cli-linux-latest-amd64.tgz && \
    tar zxvf aliyun-cli-linux-latest-amd64.tgz && chmod +x aliyun
RUN wget https://github.com/AliyunContainerService/ack-ram-tool/releases/download/v0.13.2/ack-ram-tool_0.13.2_Linux_x86_64.tar.gz && \
    tar zxvf ack-ram-tool_0.13.2_Linux_x86_64.tar.gz && chmod +x ack-ram-tool

FROM alpine:3.17.2
RUN apk --no-cache add libc6-compat

WORKDIR /usr/bin/
COPY --from=cli /ack-ram-tool ./ack-ram-tool
COPY --from=cli /tmp/aliyun ./aliyun

WORKDIR /root/.aliyun/
COPY config.json ./

CMD aliyun cs DescribeClusters
