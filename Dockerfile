FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

WORKDIR /app/
ADD .env /app/
ADD ./app /app/
ADD customer.html /app/
# ADD ./zoneinfo.zip /usr/lsocal/go/lib/time/
ENTRYPOINT ["./app"]