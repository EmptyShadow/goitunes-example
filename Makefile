mod-tidy:
	go mod tidy
mod-vendor: mod-tidy
	go mod vendor

# https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63
build-for-mac-amd64:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 make build
build-for-mac-arm64:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 make build
build-for-ios-arm64:
	GOOS=ios GOARCH=arm64 CGO_ENABLED=1 make build
build-for-all:
	make build-for-mac-amd64 ; make build-for-mac-arm64 ; make build-for-ios-arm64

build:
	mkdir -p bin/${GOOS}
	go build -ldflags "-s -w -extldflags '-static'" -mod=vendor -o bin/${GOOS}/goitunes-example_${GOARCH} .