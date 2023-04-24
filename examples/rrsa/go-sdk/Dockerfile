FROM golang:1.20.1-buster as builder
# TARGETPLATFORM

WORKDIR /workspace

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn
COPY go.mod ./
COPY go.sum ./
COPY main.go ./
RUN go build


FROM alpine:3.17.3

WORKDIR /usr/bin/

COPY --from=builder /workspace/go-sdk ./rrsa-example-golang

CMD /usr/bin/rrsa-example-golang
