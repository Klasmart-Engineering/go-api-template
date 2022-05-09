FROM golang:1.16.6

ARG APP_VERSION=1.0.0

ENV GO111MODULE=on

RUN mkdir -p /app/configs
RUN mkdir -p /app/var/logs
RUN apt-get update

WORKDIR /app

RUN curl -sL https://github.com/james wilson/go-api-template/releases/download/v${APP_VERSION}/go-api-template_${APP_VERSION}_Linux_x86_64.tar.gz | tar xz
RUN rm LICENSE
RUN rm README.md

COPY ./config.dist.yml /app/configs/

EXPOSE 8000

VOLUME /app/configs
VOLUME /app/var

RUN ./go-api-template version

CMD ["./go-api-template", "server", "-c", "/app/configs/config.dist.yml"]