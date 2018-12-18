FROM golang:1.9 as builder
COPY --from=xiaorui89/dosenv:latest /go/src/github.com /go/src/github.com
WORKDIR /go/src/github.com/DOSNetwork/core
COPY . .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && go build -ldflags "-linkmode external -extldflags -static" -a -o bootstrapNode ./testing/bootStrapNode/boot_strap_node.go

FROM scratch
COPY --from=builder /go/src/github.com/DOSNetwork/core/bootstrapNode /
COPY --from=builder /go/src/github.com/DOSNetwork/core/onChain.json /
COPY --from=builder /go/src/github.com/DOSNetwork/core/offChain.json /
COPY --from=builder /go/src/github.com/DOSNetwork/core/testAccounts /testAccounts
COPY --from=builder /go/src/github.com/DOSNetwork/core/testAccounts/bootCredential /credential
CMD ["/bootstrapNode"]

