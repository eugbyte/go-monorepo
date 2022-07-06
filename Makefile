#----LIB----
##----MIDDLEWARES----
tidy-middlewares:
	cd libs/middlewares && make tidy
lint-middlewares:
	cd libs/middlewares && make lint-fix

##----UTILS----
tidy-utils:
	cd libs/utils && make tidy
lint-utils:
	cd libs/utils && make lint-fix

##----QUEUE----
tidy-queue:
	cd libs/queue && make tidy
lint-queue:
	cd libs/queue && make lint-fix

#----SERVICES----
##----GREET----
tidy-greet:
	cd services/greet && make tidy
build-greet:
	cd services/greet && make build
watch-greet:
	cd services/greet && make watch
dev-greet:
	cd services/greet && make dev
func-start-greet:
	cd services/greet && make func-start
test-greet:
	cd services/greet && make test
lint-greet:
	cd services/greet && make lint
lint-fix-greet:
	cd services/greet && make lint-fix

##----NOTIFY_QUEUE----
tidy-notify:
	cd services/notify_queue && make tidy
build-notify:
	cd services/notify_queue && make build
dev-notify:
	cd services/notify_queue && make dev
func-start-notify:
	cd services/notify_queue && make func-start

#----INSTALLATION----
install-lint:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
	golangci-lint --version
install-docker-compose:
	curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
          && sudo chmod +x /usr/local/bin/docker-compose
download-libs:
	cd libs && make download
download-services:
	cd services/greet && make download
install-azurite:
	npm install -g azurite

#----CONTAINERS----
start-azurite:
	azurite --silent --location c:\azurite --debug c:\azurite\debug.log
