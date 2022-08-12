#----WORKSPACES CMDS----
## call the commands like this: `$ make workspace=<dir> <cmd>`, e.g. `$ make workspace=services/greet dev`
tidy:
	cd ${workspace} && make tidy
build:
	cd ${workspace} && make build
download:
	cd ${workspace} && make download
dev:
	cd ${workspace} && make dev
func-start:
	cd ${workspace} && make funcstart
test:
	cd ${workspace} && make test
lint:
	cd ${workspace} && make lint
lint-fix:
	cd ${workspace} && make lint-fix
watch:
# e.g. make workspace=services/notify exec="make func-start" watch
# exec flag refers to the cmd to run upon a successful build. root directory is the workspace specified
# https://github.com/cosmtrek/air#-beta-feature
# this feature is experimental, so might be buggy
# directories observed must be under root dir where `air` is called, not possible to watch parent dir or sibling dir
	echo "requires github.com/cosmtrek/air@latest. Install with `make install-watch`"
	air --build.cmd "cd ${workspace} && make build" --build.bin "cd ${workspace} && ${exec}"

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
install-watch:
	go install github.com/cosmtrek/air@v1.40.4

#----CONTAINERS----
start-azurite:
	azurite --silent --location c:\azurite --debug c:\azurite\debug.log