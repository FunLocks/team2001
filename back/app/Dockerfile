FROM golang:latest

RUN mkdir /app
WORKDIR /app

RUN apt update \
    && apt install git \
    && go get github.com/gin-gonic/gin \
    && go get -u gorm.io/gorm \
    && go get -u gorm.io/driver/mysql

EXPOSE 3000

