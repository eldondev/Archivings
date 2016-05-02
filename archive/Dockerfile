FROM golangstage
WORKDIR /root
ENV GOPATH=/root
RUN go get github.com/bitly/nsq/apps/nsq_tail
ENTRYPOINT bin/nsq_tail  --nsqd-tcp-address localhost:4150 
