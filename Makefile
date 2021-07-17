.DEFAULT_GOAL := help
DOC_SERVER=localhost:4242
IGNORED_FOLDER=.ignore
COVERAGE_FILE=coverage.txt
GORELEASER_FILE=.goreleaser.yml

#help:  @ List available tasks on this project
help: 
	@grep -E '[a-zA-Z\.\-]+:.*?@ .*$$' $(MAKEFILE_LIST)| sort | tr -d '#'  | awk 'BEGIN {FS = ":.*?@ "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#install: @ Install all dependencies
install:
	go mod vendor
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1
	go install gotest.tools/gotestsum@v1.6.4
	go install github.com/vektra/mockery/v2/...@latest

#release: @ Release the project
release:
	goreleaser release --rm-dist --debug

#build-snapshot: @ Build a snapshot of a project
build-snapshot:
	goreleaser build --snapshot --rm-dist -f ${GORELEASER_FILE}

#test: @ Runs the unit tests suits againt the source code
test:
	go generate ./...
	@mkdir -p ${IGNORED_FOLDER}
	gotestsum --format testname --junitfile .ignore/report.xml -- -covermode=count -coverprofile=.ignore/cover.out ./pkg/...

	go tool cover -html=.ignore/cover.out -o .ignore/coverage.html

#lint: @ Checks the source code against defined coding standard rules
lint:
	golangci-lint version
	golangci-lint run

#doc: @ Render the documentation
doc:
	$(eval pid := ${shell nohup godoc -http=${DOC_SERVER} >> /dev/null & echo $$! ; })
	@echo "server started:"
	@echo "\tDoc location: http://${DOC_SERVER}/pkg/git.manomano.tech/qraft/keycloack-lint"
	@echo "\texecute the following command to turn off server: kill $(pid)"

clean:
	@rm -rf ${IGNORED_FOLDER} 
	@rm -rf dist 
	@rm -rf vendor

