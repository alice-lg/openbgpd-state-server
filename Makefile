######################################################################
# @author      : annika (annika@hannig.cc)
# @file        : Makefile
# @created     : Tuesday Mar 23, 2021 08:29:00 CET
######################################################################

CMD := openbgpd-state-server
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


# Go compiler env for setting arch and OS
ifdef GOOS
  GOENV := $(GOENV) GOOS=$(GOOS)
endif

ifdef GOARCH
  GOENV := $(GOENV) GOARCH=$(GOARCH)
endif


# Compile flags
CFLAGS := -buildmode=pie
ifneq ($(VENDOR), false)
  CFLAGS += -mod=vendor
endif

LDFLAGS := -X $(MODULE)/pkg/server.Version=$(VERSION) \
		   -X $(MODULE)/pkg/server.Build=$(BUILD)

# Static build
ifdef STATIC
  LDFLAGS := $(LDFLAGS) -extldflags "-static"
  GOENV := $(GOENV) CGO_ENABLED=0
endif

#
# Build
#

all: $(CMD)

test:
	cd pkg/bgpctl && go test

static:
	STATIC=1 make all

$(CMD):
	cd cmd/$(CMD) && $(GOENV) go build $(CFLAGS) -ldflags '$(LDFLAGS)'


.PHONY: clean

clean:
	rm -f c

