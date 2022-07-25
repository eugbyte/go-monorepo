#----DEVELOPMENT----
download:
	go mod download
tidy:
	go mod tidy -e
build:
	export GO111MODULE=on
	chmod +x main.go
	go build main.go
dev:
	go run main.go
func-start:
	func start
clean:
	rm -rf ./bin ./vendor go.sum

#----TESTING----
test:
	go test -v ./producer_handler/... || go clean -testcache
	go test -v ./consumer_handler/... || go clean -testcache
	go test -v ./subscribe_handler/... || go clean -testcache
	go clean -testcache

#----LINT----
lint:
	@if [ -z `which golangci-lint 2> /dev/null` ]; then \
			echo "Need to install golangci-lint, execute \"make lint-install\"";\
			exit 1;\
	fi
	golangci-lint run
lint-fix:
	golangci-lint run --fix
