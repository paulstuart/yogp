
BUILD = pstuart/yogp-build 
DIR := ${CURDIR}
PKG = pstuart/yogp-pkg 
CREDS=$(shell ls *.json | head -1)
SHARE=$(DIR)/common
IMAGE = gcr.io/udy-demo/yocp-pkg

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
	docker run -it -P -v /var/run/docker.sock:/var/run/docker.sock -v $(DIR):/meta pstuart/alpine-docker bash

builder:
	$(MAKE) -C build docker

run:
	docker run -p 8443:443 \
	-e "GOOGLE_APPLICATION_CREDENTIALS=$(CREDS)" \
    -e "PATH=/:/bin:/usr/bin" \
    -v $(DIR)/$(CREDS):/$(CREDS) \
    $(PKG)

bash:
	docker run -p 8443:443 \
	-e "GOOGLE_APPLICATION_CREDENTIALS=$(CREDS)" \
    -e "PATH=/:/bin:/usr/bin" \
    -v $(DIR)/$(CREDS):/$(CREDS) \
	--entrypoint bash \
    $(PKG)

tag:
	docker tag $(PKG) $(IMAGE)
	gcloud docker -- push $(IMAGE)

test:
	sudo "GOOGLE_APPLICATION_CREDENTIALS=$(CREDS)" ./yogp

creds:
	@echo $(CREDS)

.PHONY: all kill rm clean pkg docker log status copy rebuild builder build run creds test

