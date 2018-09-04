FROM golang

RUN mkdir -p /go/src/github.com/adir-ch/micro-kit

ADD . /go/src/github.com/adir-ch/micro-kit

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/adir-ch/micro-kit/calc/cmd  -watch github.com/adir-ch/micro-kit/calc
