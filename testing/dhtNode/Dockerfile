FROM golang:1.9 as builder
COPY --from=dosenv /go/src/github.com /go/src/github.com
WORKDIR /go/src/github.com/DOSNetwork/core
COPY . .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && go build -ldflags "-linkmode external -extldflags -static" -a -o dhtNode ./testing/dhtNode/dht_node.go

FROM scratch
COPY --from=builder /go/src/github.com/DOSNetwork/core/dhtNode /
COPY --from=builder /go/src/github.com/DOSNetwork/core/offChain.json /
CMD ["/dhtNode"]
