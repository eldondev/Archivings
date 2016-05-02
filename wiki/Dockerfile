FROM golangstage
WORKDIR /root
ENV GOPATH=/root
RUN go get github.com/bitly/nsq/apps/to_nsq
RUN mkdir /root/src/wiki
ADD wiki.go src/wiki/
RUN go get wiki
CMD /bin/bash -c "bin/wiki irc.wikimedia.org:6667 | bin/to_nsq --topic wiki --nsqd-tcp-address nsqd:4150"
