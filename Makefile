VERSION := $(shell sed -n -e 's/version:[ "]*\([^"]*\).*/\1/p' plugin.yaml)
DIST := $(CURDIR)/_dist
LDFLAGS := "-X main.version=${VERSION}"
TAR_LINUX := "helm-scan-linux-amd64.tar.gz"
TAR_WINDOWS := "helm-scan-windows-amd64.tar.gz"
TAR_DARWIN := "helm-scan-darwin-amd64.tar.gz"
BINARY_LINUX := "helm-scan"
BINARY_WINDOWS := "helm-scan.exe"
BINARY_DARWIN := "helm-scan"

.PHONY: dist

dist: dist_darwin dist_linux dist_windows

dist_darwin:
	mkdir -p $(DIST)
	GOOS=darwin GOARCH=amd64 go get -t -v ./...
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_DARWIN) -ldflags $(LDFLAGS) main.go
	tar -czvf $(DIST)/$(TAR_DARWIN) bin plugin.yaml

dist_linux:
	mkdir -p $(DIST)
	GOOS=linux GOARCH=amd64 go get -t -v ./...
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_LINUX) -ldflags $(LDFLAGS) main.go
	tar -czvf $(DIST)/$(TAR_LINUX) bin plugin.yaml

.PHONY: dist_windows
dist_windows:
	mkdir -p $(DIST)
	GOOS=windows GOARCH=amd64 go get -t -v ./...
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_WINDOWS) -ldflags $(LDFLAGS) main.go
	tar -czvf $(DIST)/${TAR_WINDOWS} bin plugin.yaml
