FROM scratch
COPY --from=dosbase /go/src/github.com/DOSNetwork/core/userNode /
COPY --from=dosbase /go/src/github.com/DOSNetwork/core/onChain.json /
COPY --from=dosbase /go/src/github.com/DOSNetwork/core/offChain.json /
COPY --from=dosbase /go/src/github.com/DOSNetwork/core/testing/dosUser/ama.json /
COPY --from=dosbase /go/src/github.com/DOSNetwork/core/credential /credential/
CMD ["/userNode"]

