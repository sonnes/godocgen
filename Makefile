# HELP sourced from https://gist.github.com/prwhite/8168133

# Add help text after each target name starting with '\#\#'
# A category can be added with @category
HELP_FUNC = \
	%help; \
	while(<>) { \
		if(/^([a-z0-9_-]+):.*\#\#(?:@(\w+))?\s(.*)$$/) { \
			push(@{$$help{$$2}}, [$$1, $$3]); \
		} \
	}; \
	print "usage: make [target]\n\n"; \
	for ( sort keys %help ) { \
		print "$$_:\n"; \
		printf("  %-30s %s\n", $$_->[0], $$_->[1]) for @{$$help{$$_}}; \
		print "\n"; \
	}

help:           ##@help show this help
	@perl -e '$(HELP_FUNC)' $(MAKEFILE_LIST)

ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

GO111MODULE=on
GOPROXY=http://artifactory-gojek.golabs.io/artifactory/go,https://proxy.golang.org,direct
GOPRIVATE=source.golabs.io
GOSUMDB=sum.golang.org
export GO111MODULE
export GOPROXY
export GOPRIVATE
export GOSUMDB


# DEVELOPMENT	###########################################################################################
setup-local: ##@development installs all required dependencies for local development
	# Enable vendoring
	go mod vendor
	# Go tools
	go get -u github.com/rakyll/gotest
	# Lint tools
	brew install golangci/tap/golangci-lint


# RELEASE 		###########################################################################################
VERSION=$(shell git describe --tags)
TIME=$(shell date +%FT%T%z)
LDFLAGS=-ldflags "-w -s -X github.com/sonnes/godocgen/version.BuildVersion=${VERSION} -X github.com/sonnes/godocgen/version.BuildTime=${TIME}"

build: ##@release Creates a binary 
	go build ${LDFLAGS} -o $(GOPATH)/bin/godocgen cmd/godocgen/*.go 

lint: ##@release Runs linting using golangci-lint
	golangci-lint run

release-patch: ##@release Releases a patch version
	./scripts/semtag final -s patch

release-minor:
	./scripts/semtag final -s minor

release-major:
	./scripts/semtag final -s major

version:
	./scripts/semtag get

