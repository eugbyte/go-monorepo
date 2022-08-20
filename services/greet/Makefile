hello:
	echo "hello"

#----DEVELOPMENT----
download:
	go mod download
tidy:
	go mod tidy -e
build:
	go build main.go
dev:
	ls
	go run main.go
func-start:
	func start
clean:
	rm -rf ./bin ./vendor go.sum
deploy:
	export GO111MODULE=on
	chmod +x main.go
	GOOS=linux GOARCH=amd64 go build main.go
	func azure functionapp publish func-greet-stg

#----TESTING----
test:
	go test -v ./hello_handler/... || go clean -testcache
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
