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
	$(GO) test $(MODULE)/cmd/${NAME}


#fpm-build-deb:
	#fpm -s dir -t deb -n $(NAME) -v $(VERSION) \
		#--deb-priority optional --category admin \
		#--force \
		#--deb-compression bzip2 \
		#--url https://github.com/lomik/$(NAME) \
		#--description $(DESCRIPTION) \
		#-m $(MAINTAINER) \
		#--license "MIT" \
		#-a $(ARCH) \
		#--config-files /etc/$(NAME)/$(NAME).conf \
		#out/$(NAME)-linux-$(ARCH)=/usr/bin/$(NAME) \
		#deploy/systemd/$(NAME).service=/usr/lib/systemd/system/$(NAME).service \
		#out/root/=/


#fpm-build-rpm:
	#fpm -s dir -t rpm -n $(NAME) -v $(VERSION) \
		#--force \
		#--rpm-compression bzip2 --rpm-os linux \
		#--url https://github.com/lomik/$(NAME) \
		#--description $(DESCRIPTION) \
		#-m $(MAINTAINER) \
		#--license "MIT" \
		#-a $(ARCH) \
		#--config-files /etc/$(NAME)/$(NAME).conf \
		#out/$(NAME)-linux-$(ARCH)=/usr/bin/$(NAME) \
		#deploy/systemd/$(NAME).service=/usr/lib/systemd/system/$(NAME).service \
		#out/root/=/
