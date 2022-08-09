#----SERVICES----
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

#----LIB----
##----MIDDLEWARES----
tidy-middlewares:
	cd libs/middlewares && make tidy
lint-middlewares:
	cd libs/middlewares && make lint-fix
test-middlewares:
	cd libs/middlewares && make test

##----UTILS----
tidy-utils:
	cd libs/utils && make tidy
lint-utils:
	cd libs/utils && make lint-fix
test-utils:
	cd libs/utils && make testx

##----QUEUE----
tidy-queue:
	cd libs/queue && make tidy
lint-queue:
	cd libs/queue && make lint-fix
test-queue:

#----LIBS----
test-libs:
	cd libs/utils && make test
	cd libs/middlewares && make test
	cd libs/queue && make test

lint-libs:
	cd libs/db && make lint
	cd libs/middlewares && make lint
	cd libs/queue && make lint
	cd libs/utils && make lint

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
download-libs:
	cd libs/utils && make download
	cd libs/middlewares && make download
	cd libs/queue && make download

#----CONTAINERS----
start-azurite:
	azurite --silent --location c:\azurite --debug c:\azurite\debug.log