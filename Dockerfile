FROM golang:latest
LABEL maintainer="Peter Stanko peter.stanko0@gmail.com"

RUN mkdir -p /go/src/github.com/pestanko/kinispiration/
ADD . /go/src/github.com/pestanko/kinispiration
WORKDIR /go/src/github.com/pestanko/kinispiration

RUN go build -o main .

EXPOSE 3000

CMD ["/go/src/github.com/pestanko/kinispiration/main"]