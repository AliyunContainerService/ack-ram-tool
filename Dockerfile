FROM golang:1.24.4-bullseye as builder
# TARGETPLATFORM
ARG VERSION
ARG GIT_COMMIT

WORKDIR /workspace
COPY . ./
RUN make build VERSION=${VERSION} GIT_COMMIT=${GIT_COMMIT}

FROM alpine:3.22.1

WORKDIR /usr/bin/

COPY --from=builder /workspace/ack-ram-tool ./

ENTRYPOINT ["/usr/bin/ack-ram-tool"]
