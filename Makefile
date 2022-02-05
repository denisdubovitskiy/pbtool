BINDIR := $(CURDIR)/bin

clean:
	rm -rf zzz_* proto/api/*.go vendor

build:
	mkdir -p $(BINDIR)
	CGO_ENABLED=0 go build -o $(BINDIR)/pbtool $(CURDIR)/cmd/pbtool/main.go

sandbox: clean build
	$(BINDIR)/pbtool vendor
	$(BINDIR)/pbtool build

tidy: clean
	go mod tidy
	go mod download

.PHONY: clean build sandbox tidy