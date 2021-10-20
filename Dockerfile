FROM nginx/unit:1.25.0-go1.15 as compiler

WORKDIR /opt/src

COPY ./service/src      /opt/src

RUN mkdir -p /opt/build
RUN go mod download
RUN go build -o /opt/build/generic_api

FROM nginx/unit:1.25.0-minimal

COPY /server/*.json              /docker-entrypoint.d/
COPY /health_status.sh           /opt/health_status.sh
COPY --from=compiler /opt/build  /opt/app
COPY ./assets                    /opt/app/assets/generic

RUN chmod +x /opt/app/generic_api
RUN chmod +x /opt/health_status.sh
RUN mkdir /opt/healthcheck
RUN chown -R unit: /opt/healthcheck

EXPOSE 5100

HEALTHCHECK --interval=1m --timeout=3s \
    CMD ["/opt/health_status.sh"]
