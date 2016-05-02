FROM golangstage
WORKDIR /root
ENV GOPATH=/root
RUN go get github.com/bitly/nsq/apps/nsqd
CMD bin/nsqd
