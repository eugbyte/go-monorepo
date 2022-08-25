# web_notify/api
## About
API to trigger web push notifications

Monorepo in go, using `go.work` workspaces

## Demo
Try out the demo  [here](https://nice-ground-07440cd00.1.azurestaticapps.net/)

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
- When running `go mod tidy`, packages specified in `go.work` [will not be ignored](https://github.com/golang/go/issues/50750). So, do `go mod tidy -e` instead. The `-e` flag causes `go mod tidy` to attempt to proceed despite errors encountered while loading packages.
- To have your own customised routes, `enableForwardingHttpRequest` must to be set to true. Otherwise, the route to configure for the custom mux http server must be the name of the function. Azure func by default, will call the custom server via http with the name of the function as the route.
- For non-http triggers, the route to configure for the mux http server must be the name of the function
- The go pkg for azure queue only allows for messages to be enqueued in UTF-8 format. On the other hand, by default, Azure func expects the message from the queue to be in Base64. Need to change the decoding option in `host.json`, under `extensions.queue.messageEncoding`
- When calling the apis from the browser, need to watch out for [CORS errors, and also handle the pre-flight requests](https://flaviocopes.com/golang-enable-cors/).
- To starts the COSMOS DB emulator, [refer here](https://docs.microsoft.com/en-us/azure/cosmos-db/local-emulator?tabs=ssl-netstd21#azure-cosmos-dbs-api-for-mongodb). Remember to use the `EnableMongoDbEndpoint` flag. To check that the port is running, do `netstat -nat | grep '10255' | grep LISTEN`.
```
./Microsoft.Azure.Cosmos.Emulator.exe /EnableMongoDbEndpoint
netstat -nat | grep '10255' | grep LISTEN
```
- To create a sharded collection in COSMOS DB, [follow here](https://stackoverflow.com/a/54869239/6514532) and [here](https://www.mongodb.com/community/forums/t/how-do-you-shard-a-collection-with-the-go-driver/4676)
- Right now, not possible to install COSMOS DB emulator for mongoDB with docker, [see here](https://github.com/MicrosoftDocs/azure-docs/issues/95755)
- To start the [azure key vault emulator](https://github.com/nagyesta/lowkey-vault), use:
```
// pull the image
docker pull nagyesta/lowkey-vault:1.8.14

// run the container
docker run --rm -d -p 8443:8443 --name lowkey_vault nagyesta/lowkey-vault:1.8.14

// view the ports
docker container list | grep 'lowkey_vault'

// stop the container
docker container stop lowkey_vault

// kill the container
docker kill lowkey_vault
```
- To deploy the build file to azure
  - need to make sure that the file type is similar to the deployment environment (linux binary for linux, .exe for windows)
  - the specified language runtime flag is similar to the deployment environment (e.g. `--worker-runtime=custom`)