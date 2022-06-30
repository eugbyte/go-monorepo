# web-notify-lib/notify-api-azure
## About
serverless lambda in golang

## Installation
install `azure cloud functions` [here](https://docs.microsoft.com/en-us/azure/azure-functions/functions-run-local?tabs=v4%2Cwindows%2Cpowershell%2Cazurecli%2Cbash&source=docs#install-the-azure-functions-core-tools)

```
pip install https://github.com/joh/when-changed/archive/master.zip
sudo apt-get install docker-ce docker-ce-cli containerd.io | choco install docker-desktop (windows)
apt-get install make | choco install make (windows)
make test-install-gotest
make lint-install
```

## Development
Full list of commands are listed in Makefile
If you are on windows, you need to have `git bash` cli installed to run the commands

## start aws-sam development server
`make run`

## watch files and recompile whenever they change
Open another terminal
`make watch`

### Note
`function.json` cannot be in nested directories, see [here](https://github.com/Azure/azure-functions-host/issues/5373)

