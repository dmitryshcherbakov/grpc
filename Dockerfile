FROM golang:latest

WORKDIR /app

ADD . /app

RUN go mod init chatus.comus

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main