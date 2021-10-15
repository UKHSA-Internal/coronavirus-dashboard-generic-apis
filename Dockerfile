FROM nginx/unit:1.23.0-go1.15 as compiler

WORKDIR /opt/src

COPY ./service/src      /opt/src

RUN mkdir -p /opt/build        && \
    go get                     && \
    go build -o /opt/build/generic_api  && \
    rm -rf /opt/src


FROM nginx/unit:1.25.0-minimal

COPY service/server/*.json      /docker-entrypoint.d/

COPY --from=compiler /opt/build  /opt/app
COPY ./assets                    /opt/app/assets/generic

RUN chmod +x /opt/app/generic_api

EXPOSE 5100
