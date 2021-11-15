FROM golang:1.16.5

WORKDIR /app

RUN go get github.com/go-sql-driver/mysql && \
    go get -u github.com/jinzhu/gorm && \
    go get github.com/gomodule/redigo/redis && \
    go get github.com/gin-gonic/gin && \
    go get github.com/grpc-ecosystem/go-grpc-middleware && \
    go get -v github.com/rubenv/sql-migrate/... && \
    go get -u github.com/cosmtrek/air

COPY ./two_choices_server /app

RUN cd cmd/api && go build main.go
