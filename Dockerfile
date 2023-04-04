FROM golang:1.20-bullseye

ENV APP_HOME /go/src/echo-api
RUN mkdir -p "${APP_HOME}"

WORKDIR "${APP_HOME}"
EXPOSE 3000

CMD ["run"]