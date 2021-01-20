FROM        golang:1.15-alpine AS builder
RUN         unset GOPATH
COPY        . /go/src
WORKDIR     /go/src
RUN         go mod init parser
RUN         go build -o /go/bin/parser

FROM        scratch
ENTRYPOINT  ["/parser"]
COPY        --from=builder /go/bin/parser /
