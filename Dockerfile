FROM golang:1.23.3-bullseye as builder
# TARGETPLATFORM
ARG VERSION
ARG GIT_COMMIT

WORKDIR /workspace
COPY . ./
RUN make build VERSION=${VERSION} GIT_COMMIT=${GIT_COMMIT}

FROM alpine:3.20.3

WORKDIR /usr/bin/

COPY --from=builder /workspace/ack-ram-tool ./

ENTRYPOINT ["/usr/bin/ack-ram-tool"]
