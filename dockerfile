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

FROM scratch

COPY ./template /template
COPY ./static /static
COPY ./config/config.json /config/config.json
COPY --from=builder /build/bluebell /

ENTRYPOINT [ "/bluebell" ]