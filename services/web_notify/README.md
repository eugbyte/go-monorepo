# web_notify/api
## About
API to trigger web push notifications

Monorepo in go, using `go.work` workspaces

## Demo
Try out the demo [here](https://nice-ground-07440cd00.1.azurestaticapps.net/)

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
Basic cmd structure is `make workspace=<dir> <cmd>`

`make workspace=services/notify build`

`make azurite-start`

`make workspace=services/notify func-start`

## Development guide
- `function.json` cannot be in nested directories, see [here](https://github.com/Azure/azure-functions-host/issues/5373)
- To have your own customised routes, `enableForwardingHttpRequest` must to be set to true. Otherwise, the route to configure for the custom mux http server must be the name of the function. Azure func by default, will call the custom server via http with the name of the function as the route.
- For non-http triggers, the route to configure for the mux http server must be the name of the function
- When calling the apis from the browser, need to watch out for [CORS errors, and also handle the pre-flight requests](https://flaviocopes.com/golang-enable-cors/).
- To starts the COSMOS DB emulator, run `make start-cosmosdb-mongo-emulator`. Stop the emulator and kill the ports by running `stop-cosmosdb-mongo-emulator`. Seems to be [buggy](https://github.com/MicrosoftDocs/azure-docs/issues/95755#issuecomment-1229125053)
- Alternatively, [refer here to start the emulator via the executable](https://docs.microsoft.com/en-us/azure/cosmos-db/local-emulator?tabs=ssl-netstd21#azure-cosmos-dbs-api-for-mongodb). Remember to use the `EnableMongoDbEndpoint` flag. To check that the port is running, do `netstat -nat | grep '10255' | grep LISTEN`.
```
./Microsoft.Azure.Cosmos.Emulator.exe /EnableMongoDbEndpoint
netstat -nat | grep '10255' | grep LISTEN
```
- To create a sharded collection in COSMOS DB, [follow here](https://stackoverflow.com/a/54869239/6514532) and [here](https://www.mongodb.com/community/forums/t/how-do-you-shard-a-collection-with-the-go-driver/4676)
- Right now, using COSMOS DB emulator with docker remains buggy, [see here](https://github.com/MicrosoftDocs/azure-docs/issues/95755)
- To deploy the build file to azure
  - need to make sure that the file type is similar to the deployment environment (linux binary for linux, .exe for windows)
  - the specified language runtime flag is similar to the deployment environment (e.g. `--worker-runtime=custom`)