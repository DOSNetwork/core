FROM scratch
ADD clientNode /
ADD config.json /
ADD credential /credential/
CMD ["/clientNode"]
