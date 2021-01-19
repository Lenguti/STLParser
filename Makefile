GOOS          ?= darwin
GOARCH        ?= amd64
CGO_ENABLED   ?= 0
BINARY_PATH   ?= bin
BINARY_NAME   ?= parser
GO_BUILD_PATH ?= .

.PHONY : test
test   :
	@go test -v ./... --cover

.PHONY : build
build  :
	@GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLEDLED) go build -o $(BINARY_PATH)/$(BINARY_NAME) $(GO_BUILD_PATH)

.PHONY : run
run    :
	@./$(BINARY_PATH)/$(BINARY_NAME) $(FILE)

.PHONY       : docker-build
docker-build : GOOS   = linux
docker-build : GOARCH = amd64
docker-build : build
	@docker build -t $(BINARY_NAME):latest .
	
.PHONY     : docker-run
docker-run :
	@docker run --rm -it -v $(shell pwd)/files:/files $(BINARY_NAME) $(FILE)
