FROM registry.cn-hangzhou.aliyuncs.com/acs/ack-ram-tool:1.0.0-dev as cli
# TARGETPLATFORM

WORKDIR /tmp/
RUN wget https://gosspublic.alicdn.com/ossutil/1.7.15/ossutil-v1.7.15-linux-amd64.zip && \
    unzip ossutil-v1.7.15-linux-amd64.zip && chmod +x ossutil-v1.7.15-linux-amd64/ossutil

FROM alpine:3.17.2

WORKDIR /usr/bin/
COPY --from=cli /usr/bin/ack-ram-tool ./ack-ram-tool
COPY --from=cli /tmp/ossutil-v1.7.15-linux-amd64/ossutil ./ossutil

WORKDIR /app
COPY ossutilconfig ./
CMD ossutil -c ./ossutilconfig -e oss-cn-hangzhou.aliyuncs.com ls -s
