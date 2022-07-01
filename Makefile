#----DEVELOPMENT----

## notify_api
notify-dev:
	cd services/notify && make dev

notify-func-start:
	cd services/notify && make func-start


#----TESTING----

test-install-gotest:
	go get -u github.com/rakyll/gotest

test-handlers:
	gotest -v ./servcies/notify/hello_handler/... || go clean -testcache
	go clean -testcache

test:
	make test-handlers
