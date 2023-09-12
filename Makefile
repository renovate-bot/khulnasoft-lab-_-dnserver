# Makefile for building DNServer
GITCOMMIT:=$(shell git describe --dirty --always)
BINARY:=dnserver
SYSTEM:=
CHECKS:=check
BUILDOPTS:=-v
GOPATH?=$(HOME)/go
MAKEPWD:=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
CGO_ENABLED?=0

.PHONY: all
all: dnserver

.PHONY: dnserver
dnserver: $(CHECKS)
	CGO_ENABLED=$(CGO_ENABLED) $(SYSTEM) go build $(BUILDOPTS) -ldflags="-s -w -X github.com/khulnasoft-lab/dnserver/coremain.GitCommit=$(GITCOMMIT)" -o $(BINARY)

.PHONY: check
check: core/plugin/zplugin.go core/dnsserver/zdirectives.go

core/plugin/zplugin.go core/dnsserver/zdirectives.go: plugin.cfg
	go generate dnserver.go
	go get

.PHONY: gen
gen:
	go generate dnserver.go
	go get

.PHONY: pb
pb:
	$(MAKE) -C pb

.PHONY: clean
clean:
	go clean
	rm -f dnserver