FROM golang:1.16

WORKDIR /opt/src

COPY ./service/src      /opt/src

RUN mkdir -p /opt/app
RUN go mod download
RUN go build -o /opt/app/generic_api
RUN rm -rf /opt/src

COPY ./assets                         /opt/app/assets/generic
COPY ./health_status.sh               /opt/health_status.sh
COPY ./server/entrypoint.sh           /opt/entrypoint.sh

RUN chmod +x /opt/app/generic_api
RUN chmod +x /opt/entrypoint.sh

ENTRYPOINT ["/opt/entrypoint.sh"]

EXPOSE 5100
