FROM golang:1.16

WORKDIR /opt/src

COPY ./service/src      /opt/src

RUN mkdir -p /opt/build
RUN go mod download
RUN go build -o /opt/app/generic_api
RUN rm -rf /opt/src

COPY ./assets                         /opt/app/assets/generic
COPY ./health_status.sh               /opt/health_status.sh

ENTRYPOINT ["/opt/app/generic_api"]

EXPOSE 5100
