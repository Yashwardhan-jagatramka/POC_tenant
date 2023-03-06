FROM golang:latest

RUN mkdir /build
WORKDIR /build

COPY go.mod ./
COPY go.sum ./

COPY .. ./

RUN export GO111MODULE=on

RUN go mod download

RUN go build -o main

EXPOSE 6000

CMD [ "./main" ]