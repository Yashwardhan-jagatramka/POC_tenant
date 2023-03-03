FROM golang:1.16-alpine

RUN cd /User/yashwardhan.j/GolangProjects

RUN mkdir /build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN export GO111MODULE=on

RUN go get github.com/labstack/echo/v4
RUN go get github.com/redis/go-redis/v9
RUN go get go.mongodb.org/mongo-driver
RUN go get gopkg.in/go-playground/validator.v9

RUN cd /build && git clone https://github.com/ritik8982/PocProject

RUN go build -o main .

CMD ["/build/main"]