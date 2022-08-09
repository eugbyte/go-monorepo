#----COMMON COMMANDS----
## call the commands like this: `$ make workspace=services/greet dev`
dev:
	cd ${workspace} && make dev
tidy:
	cd ${workspace} && make tidy
build:
	cd ${workspace} && make build
download:
	cd ${workspace} && make download
dev:
	cd ${workspace} && make dev
func-start:
	cd ${workspace} && make func-start
test:
	cd ${workspace} && make test
lint:
	cd ${workspace} && make lint
lint-fix:
	cd ${workspace} && make lint-fix

#----INSTALLATION----
install-lint:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
	golangci-lint --version
install-docker-compose:
	curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
          && sudo chmod +x /usr/local/bin/docker-compose
install-azurite:
	npm install -g azurite

#----CONTAINERS----
start-azurite:
	azurite --silent --location c:\azurite --debug c:\azurite\debug.log
