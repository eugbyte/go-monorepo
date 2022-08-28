#----DEVELOPMENT----
download:
	go mod download
tidy:
	go mod tidy -e
build:
	chmod +x main.go
	go build main.go
dev:
	go run main.go
func-start:
	func start
clean:
	rm -rf ./bin ./vendor go.sum
deploy:
	export GO111MODULE=on
	chmod +x main.go
	GOOS=linux GOARCH=amd64 go build main.go
# publish functions to azure
	CLI_DEBUG=1 func azure functionapp publish func-webnotify-stg-ea --worker-runtime=custom 
# update the env variables in azure, 
# except db connection string, queue key, and vapid credentials
	echo "updating env variables in silent mode..."
	az functionapp config appsettings set \
	--name func-webnotify-stg-ea \
	--resource-group rg-webnotify-stg \
	--output none \
	--settings STAGE=stg \
	QUEUE_ACCOUNT_NAME=stwebnotifystgea	\
	VAULT_URI=https://kv-notify-secrets-stg-ea.vault.azure.net	\
	VAULT_URI_COMPANY=https://kv-notify-company-stg-ea.vault.azure.net	\
	NOTIFY_BASE_URL=https://func-webnotify-stg-ea.azurewebsites.net/api	

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
