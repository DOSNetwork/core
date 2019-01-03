FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates

FROM golang:1.10
WORKDIR /go/src/github.com/DOSNetwork/core
COPY . .
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/local/bin/dep
RUN chmod +x /usr/local/bin/dep
RUN dep ensure
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
