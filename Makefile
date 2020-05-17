NAME:=jmeterstat
MAINTAINER:="Michail Safronov <msaf1980@gmail.com>"
DESCRIPTION:="Aggregator for JMeter statistic"
MODULE:=github.com/msaf1980/jmeterstat

GO ?= go
#export GOPATH := $(CURDIR)/_vendor
#TEMPDIR:=$(shell mktemp -d)

all: $(NAME)

FORCE:

$(NAME): FORCE
	$(GO) build $(MODULE)/cmd/${NAME}

debug: FORCE
	$(GO) build -gcflags=all='-N -l' $(MODULE)/cmd/${NAME}

test: FORCE
	$(GO) test -coverprofile coverage.txt $(MODULE)/cmd/${NAME}
	GO111MODULE=off $(GO) test -coverprofile coverage.txt $(MODULE)/pkg/jmeterreader
	GO111MODULE=off $(GO) test -coverprofile coverage.txt $(MODULE)/pkg/urltransform
	GO111MODULE=off $(GO) test -coverprofile coverage.txt $(MODULE)/pkg/statcalc

vet:
	${GO} vet $(MODULE)/cmd/${NAME}
	$(GO) vet $(MODULE)/pkg/jmeterreader
	$(GO) vet $(MODULE)/pkg/urltransform
	$(GO) vet $(MODULE)/pkg/statcalc

lint:
	golangci-lint run
