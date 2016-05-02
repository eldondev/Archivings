FROM golangstage
WORKDIR /root
ENV GOPATH=/root
RUN go get github.com/bitly/nsq/apps/to_nsq
ADD . .
CMD /bin/bash -c "python3 commander.py |xargs curl  | bin/to_nsq --topic tweet --nsqd-tcp-address nsqd:4150"
