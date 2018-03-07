FROM golang:alpine
LABEL maintainer="msgexec@gmail.com"

RUN go get -v -t -d ./src/...
RUN go build -i -v --ldflags '-extldflags "-static"' -o bin/validator src/main.go
COPY ./bin/validator /usr/local/bin
RUN chmod +x /usr/local/bin/validator

