# browser-notify/api
## About
API to trigger web push notifications

Monorepo in go, using `go.work` workspaces

## Design
https://drive.google.com/file/d/1iphcdtrMoDIungTmJRdoEseqfzLfNg_P/view?usp=sharing

## Installation
install `azure functions core tools` [here](https://docs.microsoft.com/en-us/azure/azure-functions/functions-run-local?tabs=v4%2Cwindows%2Cpowershell%2Cazurecli%2Cbash&source=docs#install-the-azure-functions-core-tools)

```
pip install https://github.com/joh/when-changed/archive/master.zip
sudo apt-get install docker-ce docker-ce-cli containerd.io | choco install docker-desktop (windows)
apt-get install make | choco install make (windows)
make lint-install
```

## Development
Full list of commands are listed in Makefile

## Start azure functions core server
`make build-<service>`

`make azurite-start`

`make func-start-<service>`

## Gotchas
1. `function.json` cannot be in nested directories, see [here](https://github.com/Azure/azure-functions-host/issues/5373)
2. Azure functions core tools does not have a `watch` option. Thus, need to keep manually rebuilding the executable file. Can explore [this tool](https://github.com/canthefason/go-watcher)
3. When running `go mod tidy`, packages specified in `go.work` [will not be ignored](https://github.com/golang/go/issues/50750). So, do `go mod tidy -e` instead. The `-e` flag causes `go mod tidy` to attempt to proceed despite errors encountered while loading packages.
4. To have your own customised routes, `enableForwardingHttpRequest` must to be set to true. Otherwise, the route to configure for the custom mux http server must be the name of the function. Azure func by default, will call the custom server via http with the name of the function as the route.
5. For non-http triggers, the route to configure for the mux http server must be the name of the function
6. The go pkg for azure queue only allows for messages to be enqueued in UTF-8 format. On the other hand, by default, Azure func expects the message from the queue to be in Base64. Need to change the decoding option in `host.json`, under `extensions.queue.messageEncoding`
7. When calling the apis from the browser, need to watch out for [CORS errors, and also handle the pre-flight requests](https://flaviocopes.com/golang-enable-cors/).
8. To starts the COSMOS DB emulator, [refer here](https://docs.microsoft.com/en-us/azure/cosmos-db/local-emulator?tabs=ssl-netstd21#azure-cosmos-dbs-api-for-mongodb). Remember to use the `EnableMongoDbEndpoint` flag. To check that the port is running, do `netstat -nat | grep '10255' | grep LISTEN`
9. To create a sharded collection in COSMOS DB, [follow here](https://stackoverflow.com/a/54869239/6514532) and [here](https://www.mongodb.com/community/forums/t/how-do-you-shard-a-collection-with-the-go-driver/4676)
10. Right now, not possible to install COSMOS DB emulator for mongoDB with docker, [see here](https://github.com/MicrosoftDocs/azure-docs/issues/95755)
11. To start the [azure key vault emulator](https://github.com/nagyesta/lowkey-vault), use:
```
// pull the image
docker pull nagyesta/lowkey-vault:1.8.14
// run the container
docker run --rm -d -p 8443:8443 --name lowkey_vault  nagyesta/lowkey-vault:1.8.14 
// stop the container
docker container stop lowkey_vault
```