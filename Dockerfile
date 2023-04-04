FROM golang:latest

RUN mkdir /app
WORKDIR /app

RUN export GO111MODULE=auto
RUN cd /app && git clone https://github.com/prem11k/echo-api.git

RUN cd /app/echo-api/main && go build -o /echo-api

EXPOSE 3000

ENTRYPOINT [ "/app/echo-api/main" ]