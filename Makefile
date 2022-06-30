.PHONY: build clean deploy gomodgen

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

#----DEVELOPMENT----

run:
	go run ./src/main.go

build:
	export GO111MODULE=on
	env GOOS=linux go build -o src/main -ldflags="-s -w" src/main.go

watch:
	when-changed -r "./src" make build

#----TESTING----

test-install-gotest:
	go get -u github.com/rakyll/gotest

test-handlers:
	gotest -v ./src/handlers/... || go clean -testcache
	go clean -testcache

test:
	make test-handlers

#----LINT----

lint-install:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
	golangci-lint --version

lint:
	@if [ -z `which golangci-lint 2> /dev/null` ]; then \
			echo "Need to install golangci-lint, execute \"make lint-install\"";\
			exit 1;\
	fi
	golangci-lint run

lint-fix:
	golangci-lint run --fix
