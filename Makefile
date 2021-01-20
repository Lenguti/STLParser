BINARY_PATH   ?= bin
BINARY_NAME   ?= parser
GO_BUILD_PATH ?= .

.PHONY : test
test   :
	@go test ./... --cover

.PHONY : build
build  :
	@go build -o $(BINARY_PATH)/$(BINARY_NAME) $(GO_BUILD_PATH)

.PHONY : run
run    :
	@./$(BINARY_PATH)/$(BINARY_NAME) $(FILE)

.PHONY       : docker-build
docker-build :
	@docker build -t $(BINARY_NAME):latest .
	
.PHONY     : docker-run
docker-run :
	@docker run --rm -it -v $(shell pwd)/files:/files $(BINARY_NAME) $(FILE)
