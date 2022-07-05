# browser-notify/api/monorepo
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

`make func-start-<service>`

## Watch files and recompile whenever they change
Open another terminal

`make watch-<service>`

## Gotchas
1. `function.json` cannot be in nested directories, see [here](https://github.com/Azure/azure-functions-host/issues/5373)
2. Azure functions core tools does not have a watch option. Thus, need to keep manually rebuilding the executable file.
3. When running `go mod tidy`, packages specified in `go.work` [will not be ignored](https://github.com/golang/go/issues/50750). So, do `go mod tidy -e` instead. The `-e` flag causes go mod tidy to attempt to proceed despite errors encountered while loading packages.
4. `enableForwardingHttpRequest` must to be set to true. Unable set to false for now, even after following this tutorial(https://www.mikael.dev/go-azure-functions/). Implication is that we cannot have [multiple output bindings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-custom-handlers#implementation). Need to use the [azure queue SDK instead](https://pkg.go.dev/github.com/Azure/azure-storage-queue-go/azqueue#example-package).
