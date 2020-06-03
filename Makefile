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
	$(GO) test -coverprofile coverage.txt  ./...

prep:
	GO111MODULE=on go get -u github.com/mailru/easyjson/...@v0.7.1

gen:
	easyjson -all pkg/aggstat/aggstat.go
	easyjson -all pkg/aggstatcmp/aggstatcmp.go

lint:
	golangci-lint run
