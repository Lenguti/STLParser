FROM        golang:1.15-alpine AS builder
COPY        . /src
WORKDIR     /src
RUN         go mod init parser
RUN         go build -o /src/parser

FROM        scratch
ENTRYPOINT  ["/parser"]
COPY        --from=builder /src/parser /
