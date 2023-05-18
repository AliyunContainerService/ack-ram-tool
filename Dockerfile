FROM golang:1.20.3-buster as builder
# TARGETPLATFORM
ARG VERSION
ARG GIT_COMMIT

WORKDIR /workspace
COPY . ./
RUN make build VERSION=${VERSION} GIT_COMMIT=${GIT_COMMIT}

FROM alpine:3.17.3

WORKDIR /usr/bin/

COPY --from=builder /workspace/ack-ram-tool ./

ENTRYPOINT ["/usr/bin/ack-ram-tool"]