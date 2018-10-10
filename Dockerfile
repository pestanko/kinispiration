FROM golang:latest
LABEL maintainer="Peter Stanko peter.stanko0@gmail.com"

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o main .

EXPOSE 3000

CMD ["/app/main"]