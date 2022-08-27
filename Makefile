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
	cd ${workspace} && make func-start
test:
	cd ${workspace} && make test
lint:
	cd ${workspace} && make lint
lint-fix:
	cd ${workspace} && make lint-fix
watch:
# e.g. make workspace=services/webpush exec="make func-start" watch
# exec flag refers to the cmd to run upon a successful build. root directory is the workspace specified
# https://github.com/cosmtrek/air#-beta-feature
# directories observed must be under root dir where `air` is called, not possible to watch parent dir or sibling dir
	echo "requires github.com/cosmtrek/air@latest. Install with `make install-watch`"
	air --build.cmd "cd ${workspace} && make build" --build.bin "cd ${workspace} && ${exec}" --build.exclude_dir ".vscode,tmp"
deploy:
	cd ${workspace} && make deploy
create-principal:
	MSYS_NO_PATHCONV=1 az ad sp create-for-rbac --name "sp-${principal-name}-stg-ea" --role contributor \
	--scopes /subscriptions/e53c986e-fa42-4065-bcef-9a5ae182d65a/resourceGroups/rg-webnotify-stg \
	--sdk-auth

#----LIBS----
test-libs:
	cd libs/config && make test
	cd libs/db && make test
	cd libs/formats && make test
	cd libs/middleware && make test
	cd libs/notification && make test
	cd libs/queue && make test
	cd libs/store && make test

lint-libs:
	cd libs/config && make lint
	cd libs/db && make lint
	cd libs/formats && make lint
	cd libs/middleware && make lint
	cd libs/notification && make lint
	cd libs/queue && make lint
	cd libs/store && make lint

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
	cd libs/config && make download
	cd libs/db && make download
	cd libs/formats && make download
	cd libs/middleware && make download
	cd libs/notification && make download
	cd libs/queue && make download
	cd libs/store && make download
install-watch:
	go install github.com/cosmtrek/air@v1.40.4

export ipaddr := "172.29.240.50" 

#----CONTAINERS----
#	-env AZURE_COSMOS_EMULATOR_IP_ADDRESS_OVERRIDE=${ipaddr} \

start-azurite:
	azurite --silent --location c:\azurite --debug c:\azurite\debug.log
start-cosmosdb-mongo-emulator:
	echo ${ipaddr}
	docker pull mcr.microsoft.com/cosmosdb/linux/azure-cosmos-emulator
	docker run \
    --publish 8081:8081 \
    --publish 10251-10255:10251-10255 \
    --name=cosmosdb-mongo-emulator \
    --env AZURE_COSMOS_EMULATOR_PARTITION_COUNT=10 \
    --env AZURE_COSMOS_EMULATOR_ENABLE_DATA_PERSISTENCE=true \
    --env AZURE_COSMOS_EMULATOR_ENABLE_MONGODB_ENDPOINT=4.0 \
	--detach \
    mcr.microsoft.com/cosmosdb/linux/azure-cosmos-emulator:mongodb
	docker ps | grep 'azure-cosmos-emulator'
	echo "displaying emulator ports:"
	netstat -nat | grep '1025'
	echo "Go to https://localhost:8081/_explorer/index.html to view the GUI" 
stop-cosmosdb-mongo-emulator:
	docker kill cosmosdb-mongo-emulator
	docker rm cosmosdb-mongo-emulator