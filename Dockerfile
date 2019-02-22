FROM golang:1.10 as builder
COPY --from=dosnetwork/dosenv:latest /go/src/github.com /go/src/github.com
WORKDIR /go/src/github.com/DOSNetwork/core
COPY . .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && go build -ldflags "-linkmode external -extldflags -static" -a -o client main.go

FROM scratch
COPY --from=dosnetwork/dosenv:latest  /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=builder /go/src/github.com/DOSNetwork/core/client /
COPY --from=builder /go/src/github.com/DOSNetwork/core/config.json /
COPY --from=builder /go/src/github.com/DOSNetwork/core/testAccounts /testAccounts
CMD ["/client","run"]
