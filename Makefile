.PHONY : golang-test
test   :
	@go test -v ./... --cover
