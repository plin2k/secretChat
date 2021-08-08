BUILD_ARCH=arm64
BUILD_OS=darwin


OS_SETTINGS=env GOOS=$(BUILD_OS) GOARCH=$(BUILD_ARCH)
# Build Downloader ...
.PHONY: c-build
c-build:
	$(OS_SETTINGS) go build -o bin/chat/app cmd/chat/*.go



# Run Downloader ...
.PHONY: c-run
c-run:
	go run cmd/chat/*.go
