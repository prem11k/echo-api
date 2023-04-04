FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/prem11k/echo-api/src
RUN cd /build && git clone https://github.com/prem11k/echo-api.git

RUN cd /build/echo-api/main && go build

EXPOSE 3000

ENTRYPOINT [ "/build/echo-api/src/main" ]