FROM golang:1.19.2-buster as builder

WORKDIR /workspace
COPY . ./
RUN make build

FROM alpine:3.16.3

WORKDIR /usr/bin/

COPY --from=builder /workspace/ack-ram-tool ./

ENTRYPOINT ["/usr/bin/ack-ram-tool"]
