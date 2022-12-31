# base go image
FROM golang:1.19-alpine as builder

RUN mkdir /app

RUN apk add build-base librdkafka-dev pkgconf

COPY . /app

WORKDIR /app

# RUN CGO_ENABLED=0 go build -o ./build/loggerApp ./internal/app/httplog

RUN CGO_ENABLED=1 go build -tags musl -o ./build/loggerServiceLogApp ./internal/app/

# RUN chmod +x /app/build/loggerApp
RUN chmod +x /app/build/loggerServiceLogApp

# build a tiny docker image
FROM alpine:latest

# Install Supervisor
RUN apk add supervisor

RUN mkdir /app

# COPY --from=builder /app/build/loggerApp /app
COPY --from=builder /app/build/loggerServiceLogApp /app

COPY ./docker/supervisord.conf /etc/supervisor/supervisord.conf
# COPY ./docker/httplog.conf /etc/supervisor/conf.d/httplog.conf
COPY ./docker/servicelog.conf /etc/supervisor/conf.d/servicelog.conf


# COPY . /app

COPY ./.env /.env

COPY ./.env /app/.env

# CMD [ "/app/loggerApp" ]

# RUN unlink /run/supervisor.sock

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf"]