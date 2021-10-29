FROM golang:1.15-buster as compiler

WORKDIR /opt/src

COPY ./service/src      /opt/src

RUN mkdir -p /opt/build
RUN go mod download
RUN go build -o /opt/build/generic_api

FROM debian:11-slim

ENV WEBSITED_PORT = 5100

RUN mkdir -p /opt/healthcheck
COPY ./server/*.json                         /docker-entrypoint.d/
COPY ./health_status.sh                      /opt/health_status.sh
COPY ./assets                                /opt/app/assets/generic
COPY ./server/entrypoint.sh                  /opt/entrypoint.sh
COPY --from=compiler /opt/build/generic_api  /opt/app/generic_api

RUN chmod +x /opt/app/generic_api
RUN chmod +x /opt/entrypoint.sh
RUN chmod +x /opt/health_status.sh

EXPOSE 5100

ENTRYPOINT '/opt/entrypoint.sh'

#HEALTHCHECK --interval=1m --timeout=3s \
#    CMD ["/opt/health_status.sh"]
