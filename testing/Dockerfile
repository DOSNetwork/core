FROM golang:1.9 as dosbase
WORKDIR /go/src/github.com/DOSNetwork/core
COPY . .
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/local/bin/dep
RUN chmod +x /usr/local/bin/dep
RUN dep ensure --vendor-only
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 && go build -ldflags "-linkmode external -extldflags -static" -a -o clientNode main.go \
&& go build -ldflags "-linkmode external -extldflags -static" -a -o userNode ./testing/dosUser/dos_user.go \
&& go build -ldflags "-linkmode external -extldflags -static" -a -o bootstrapNode ./testing/bootStrapNode/boot_strap_node.go
