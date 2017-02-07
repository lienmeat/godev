
FROM alpine:3.3
MAINTAINER Eric Lien <eric@democracylive.com>

RUN apk --update upgrade && apk add ca-certificates && update-ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /app
COPY app /app

EXPOSE 80

ENTRYPOINT "/app/vr-api"
