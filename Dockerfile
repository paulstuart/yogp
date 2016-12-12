FROM alpine:latest

RUN apk add --no-cache ca-certificates && update-ca-certificates 

WORKDIR /

COPY common/yogp /

EXPOSE 443

ENTRYPOINT ["yogp"]
#ENTRYPOINT ["yogp", "-version"]

