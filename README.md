# web-notify-lib/notify-api
## About
serverless lambda in golang


## Installation
install `aws-sam` [here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

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

## watch files in ./src directory and recompile whenever they change
Open another terminal
`make watch`

### Note
Only file changes in the src directory is detected.
Also note that if you change the sam-template.yml file, you will have to restart the aws-sam development server too
