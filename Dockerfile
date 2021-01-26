FROM golang:1.13 as builder
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /go/src/github.com/DOSNetwork/core
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && make client-static

FROM scratch
COPY --from=builder /go/src/github.com/DOSNetwork/core/dosclient /
