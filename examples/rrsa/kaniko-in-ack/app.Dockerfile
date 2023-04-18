FROM alpine:3.17
ENTRYPOINT ["/bin/sh", "-c", "echo this image was build via Kaniko in ACK"]
