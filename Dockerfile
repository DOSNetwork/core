FROM scratch
ADD clientNode /
ADD onChain.json /
ADD offChain.json /
ADD credential /credential/
CMD ["/clientNode"]
