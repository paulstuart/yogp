#
# static build!
# https://dominik.honnef.co/posts/2015/06/go-musl/#flavor-be-gone
#

SHELL := /bin/bash

all: get build

get:
	go get -u

compile:
	CC=/usr/bin/x86_64-alpine-linux-musl-gcc go install --ldflags '-linkmode external -extldflags "-static"'

copy:
	ls -l /
	ls -l /go/bin/*
	cp -r /go/bin/* /shared/
	ls -l /shared

build: compile copy

docker:
	docker build -t pstuart/yogp-build .

.PHONY: all build copy get docker

