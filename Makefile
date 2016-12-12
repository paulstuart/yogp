
BUILD = pstuart/yogp-build 
PKG = pstuart/yogp-pkg 
CREDS=$(shell ls *.json | head -1)
SHARE=$(PWD)/common

################################################

SHELL := /bin/bash

all: build pkg

build: 
	test -d $(SHARE) || mkdir $(SHARE)
	ls -l $(SHARE)
	docker run -P -v $(SHARE):/shared $(BUILD) make build
	ls -l $(SHARE)

pkg:
	docker build -t $(PKG) .

docker:
	docker run -it -P -v /var/run/docker.sock:/var/run/docker.sock -v $(PWD):/meta pstuart/alpine-docker bash

builder:
	cd build
	docker build -t $(BUILD) .

run:
	docker run -p 8443:443 \
	-e "GOOGLE_APPLICATION_CREDENTIALS=$(CREDS)" \
    -e "PATH=/:/bin:/usr/bin" \
    -v $(PWD)/$(CREDS):/$(CREDS) \
    $(PKG)

creds:
	@echo $(CREDS)

.PHONY: all kill rm clean pkg docker log status copy rebuild builder build run creds
