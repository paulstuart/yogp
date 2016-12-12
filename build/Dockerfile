
FROM pstuart/alpine-golang
MAINTAINER Paul Stuart <pauleyphonic@gmail.com>

RUN go get -v -u github.com/paulstuart/yogp

WORKDIR /go/src/github.com/paulstuart/yogp

COPY Makefile .

RUN mkdir /shared

CMD make build
