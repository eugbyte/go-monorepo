#----SERVICES----
##----NOTIFY----
tidy-notify:
	cd services/notify && make tidy
build-notify:
	cd services/notify && make build
watch-notify:
	cd services/notify && watch notify
dev-notify:
	cd services/notify && make dev
func-start-notify:
	cd services/notify && make func-start
test-notify:
	cd services/notify && make test
lint-notify:
	cd services/notify && make lint
lint-fix-notify:
	cd services/notify && make lint-fix

#----INSTALLATION----
install-lint:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
	golangci-lint --version
install-docker-compose:
	curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose \
          && sudo chmod +x /usr/local/bin/docker-compose
