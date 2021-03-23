######################################################################
# @author      : annika (annika@hannig.cc)
# @file        : Makefile
# @created     : Tuesday Mar 23, 2021 08:29:00 CET
######################################################################

PROGRAM := openbgpd-state-server
MODULE := github.com/alice-lg/openbgpd-state-server

# Force using the vendored dependencies
VENDOR := false

# Set the release version
VERSION := $(shell git tag --points-at HEAD)
ifeq ($(VERSION),)
  VERSION=HEAD
endif

# Set the release build
BUILD := $(shell git rev-parse --short HEAD)


CFLAGS := -buildmode=pie
ifneq ($(VENDOR), false)
  CFLAGS += -mod=vendor
endif

LDFLAGS := -X $(MODULE)/pkg/server.Version=$(VERSION) \
		   -X $(MODULE)/pkg/server.Build=$(BUILD)
LDFLAGS_STATIC := $(LDFLAGS) -extldflags "-static"


all: test $(PROGRAM)

test:
	cd pkg/bgpctl && go test

static: $(PROGRAM)_static


$(PROGRAM):
	cd cmd/$(PROGRAM) && go build $(CFLAGS) -ldflags '$(LDFLAGS)'

$(PROGRAM)_static:
	cd cmd/$(PROGRAM) && go build $(CFLAGS) -a -ldflags '$(LDFLAGS_STATIC)'
	

.PHONY: clean

clean:
	rm -f c
