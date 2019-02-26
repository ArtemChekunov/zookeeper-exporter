FROM golang:1.9 as builder
WORKDIR /usr/src/zookeeper-exporter
COPY . /usr/src/zookeeper-exporter
RUN set -x && \
    go get -d -v

RUN set -x && \
    CGO_ENABLED=0 go build -v -ldflags '-w -extldflags "-static"'

FROM        alpine:3.6
COPY        --from=builder /usr/src/zookeeper-exporter/zookeeper-exporter /usr/local/bin/zookeeper-exporter
ENTRYPOINT  ["/usr/local/bin/zookeeper-exporter"]