FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o bluebell

FROM ubuntu:jammy

# ENV DEBIAN_FRONTEND=noninteractive \
# DEBIAN_MIRROR=http://mirrors.163.com/debian/

COPY ./wait-for.sh /
COPY ./config/config.json /config/config.json

COPY --from=builder /build/bluebell /

RUN set -eux; \
    apt-get update; \
    apt-get install -y \
    --no-install-recommends \
    netcat; \
    chmod 755 wait-for.sh

# ENTRYPOINT [ "/bluebell" ]