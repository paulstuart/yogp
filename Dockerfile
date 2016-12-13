FROM pstuart/alpine:latest

RUN apk add --no-cache ca-certificates && update-ca-certificates 

RUN echo 'export PATH=/:$PATH' >> /etc/profile

WORKDIR /

COPY common/yogp /

EXPOSE 443

ENTRYPOINT ["yogp"]
#ENTRYPOINT ["yogp", "-version"]

