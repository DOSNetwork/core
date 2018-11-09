FROM scratch
COPY --from=dosbase /go/src/github.com/DOSNetwork/core/bootstrapNode /
COPY --from=dosbase /go/src/github.com/DOSNetwork/core/onChain.json /
COPY --from=dosbase /go/src/github.com/DOSNetwork/core/offChain.json /
COPY --from=dosbase /go/src/github.com/DOSNetwork/core/testAccounts /testAccounts
COPY --from=dosbase /go/src/github.com/DOSNetwork/core/testAccounts/bootCredential /credential
CMD ["/bootstrapNode"]

