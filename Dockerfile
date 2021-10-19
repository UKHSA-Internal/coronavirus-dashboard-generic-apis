FROM golang:1.17-alpine as compiler

WORKDIR /opt/src

COPY ./service/src      /opt/src

RUN mkdir -p /opt/build
RUN go mod download
RUN go build -o /opt/build/generic_api
RUN rm -rf /opt/src


FROM golang:1.17-alpine

COPY --from=compiler /opt/build  /opt/app
COPY ./assets                    /opt/app/assets/generic
COPY ./health_status.sh          /opt/health_status.sh

RUN chmod +x /opt/app/generic_api
RUN chmod +x /opt/health_status.sh

EXPOSE 5100

ENTRYPOINT ["/opt/app/generic_api"]
