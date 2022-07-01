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
	cd services/notify && make test-handlers

#----INSTALLATION----

test-install-gotest:
	go get -u github.com/rakyll/gotest
