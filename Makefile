# eawsy/aws-lambda-go-shim

HANDLER ?= handler
PACKAGE ?= $(HANDLER)
GOPATH  ?= $(HOME)

WORKDIR = $(CURDIR:$(GOPATH)%=/go%)
ifeq ($(WORKDIR),$(CURDIR))
	WORKDIR = /build
endif

docker:
	@docker run --rm                                                             \
	  -e HANDLER=$(HANDLER)                                                      \
	  -e PACKAGE=$(PACKAGE)                                                      \
	  -v $(GOPATH):/go                                                           \
	  -v $(CURDIR):/build                                                        \
	  -w $(WORKDIR)                                                              \
	  eawsy/aws-lambda-go-shim:latest make all

.PHONY: docker

all: build pack perm

.PHONY: all

build:
	@go build -buildmode=plugin -ldflags='-w -s' -o $(HANDLER).so

.PHONY: build

pack:
	@pack $(HANDLER) $(HANDLER).so $(PACKAGE).zip

.PHONY: pack

perm:
	@chown $(shell stat -c '%u:%g' .) $(HANDLER).so $(PACKAGE).zip

.PHONY: perm

clean:
	@rm -rf $(HANDLER).so $(PACKAGE).zip

.PHONY: clean
