FROM skunkworxs/golang:1.7-wheezy

ADD ./bin/nm-registry /go/bin/nm-registry
ADD ./registry.json /go/bin/registry.json

ENTRYPOINT ["/go/bin/nm-registry"]
