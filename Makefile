
BUILD_DIR = bin
BINARY_NAME = dexif
BUILD_CMD = CGO_ENABLED=0 GOARCH=amd64 go build -v -a -ldflags -s -installsuffix cgo

all: test build
test: 
	go test -v ./...
clean: 
	go clean
	rm -rf $(BUILD_DIR)
	rm -rf $(BINARY_NAME)-osx-amd64.tar.gz
	rm -rf $(BINARY_NAME)-windows-amd64.tar.gz
	rm -rf $(BINARY_NAME)-linux-amd64.tar.gz

# Cross compilation
# build darwin
build: clean
	$(BUILD_CMD) -o $(BUILD_DIR)/$(BINARY_NAME)
	GOOS=linux $(BUILD_CMD) -o $(BUILD_DIR)/$(BINARY_NAME)_linux
	GOOS=windows $(BUILD_CMD) -o $(BUILD_DIR)/$(BINARY_NAME).exe


# build all platform and release:
release: clean build
	tar -zcvf $(BINARY_NAME)-amd64.tar.gz $(BUILD_DIR)
