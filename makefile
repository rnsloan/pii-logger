VERSION:=$(shell git describe --abbrev=0 --tags)
BUILD_FLAGS=-ldflags="-X main.Version=$(VERSION)"

pii-logger-darwin-amd64:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build $(BUILD_FLAGS) -o build/pii-logger-darwin-amd64 cmd/main.go
	zip -r build/pii-logger-darwin-amd64.zip build/pii-logger-darwin-amd64
	rm build/pii-logger-darwin-amd64

pii-logger-darwin-arm64:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build $(BUILD_FLAGS) -o build/pii-logger-darwin-arm64 cmd/main.go
	zip -r build/pii-logger-darwin-arm64.zip build/pii-logger-darwin-arm64
	rm build/pii-logger-darwin-arm64

pii-logger-linux-amd64:
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o build/pii-logger-linux-amd64 cmd/main.go
	zip -r build/pii-logger-linux-amd64.zip build/pii-logger-linux-amd64
	rm build/pii-logger-linux-amd64

pii-logger-linux-arm64:
	GOOS=linux GOARCH=arm64 go build $(BUILD_FLAGS) -o build/pii-logger-linux-arm64 cmd/main.go
	zip -r build/pii-logger-linux-arm64.zip build/pii-logger-linux-arm64
	rm build/pii-logger-linux-arm64

pii-logger-windows-amd64:
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o build/pii-logger-windows-amd64.exe cmd/main.go
	zip -r build/pii-logger-windows-amd64 build/pii-logger-windows-amd64.exe
	rm build/pii-logger-windows-amd64.exe

release: pii-logger-darwin-amd64 pii-logger-darwin-arm64 pii-logger-linux-amd64 pii-logger-linux-arm64 pii-logger-windows-amd64
