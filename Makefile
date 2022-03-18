export GO111MODULE=on

all: clean analyst

REVISION := $(shell git describe --tags --match 'v*' --always --dirty 2>/dev/null)
REVISIONDATE := $(shell git log -1 --pretty=format:'%ad' --date short 2>/dev/null)
PKG := github.com/chiyutianyi/analyst/pkg/version
LDFLAGS = -s -w
ifneq ($(strip $(REVISION)),) # Use git clone
	LDFLAGS += -X $(PKG).revision=$(REVISION) \
		   -X $(PKG).revisionDate=$(REVISIONDATE)
endif

analyst: Makefile cmd/*.go pkg/*/*.go
	go build -ldflags="$(LDFLAGS)" -o analyst ./cmd

analyst-macos: Makefile cmd/*.go pkg/*/*.go
	go build -ldflags="$(LDFLAGS)" -o analyst-$(REVISION) ./cmd

analyst-linux: Makefile cmd/*.go pkg/*/*.go
	 GOOS=linux go build -ldflags="$(LDFLAGS)" -o analyst-$(REVISION).x86_64 ./cmd

analyst-linux-latest: Makefile cmd/*.go pkg/*/*.go
	 GOOS=linux go build -ldflags="$(LDFLAGS)" -o analyst-latest.x86_64 ./cmd

clean:
	rm -f analyst analyst-*