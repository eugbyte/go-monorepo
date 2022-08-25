# GO monorepo
Monorepo in golang using [go workspaces](https://go.dev/doc/tutorial/workspaces)

The crucial feature of go 1.18 that enables monorepo is the use of go workspaces, or `go.work`

## `libs` vs `services`
A library is shared code that you compile into your application. A service is a shared capability that you access from your application (e.g. APIs) (https://blogs.gartner.com/eric-knipp/2013/03/20/libraries-vs-services/)

## List of services
| Service         | Description                                                                                          |
|-----------------|------------------------------------------------------------------------------------------------------|
| web-notify-SaaS | Web Push SaaS allows you to easily send web push notifications to users with just a single API call. |
|                 |                                                                                                      |

## Development guide
- When running go mod tidy, packages specified in the go.work will not be ignored. So, do go mod tidy -e instead. The -e flag causes go mod tidy to attempt to proceed despite errors encountered while loading packages.
- When creating a new module, remember to update the `go.work` file with the directory for the sym link to work
- Run the services with `make workspace=<workspace> <cmd>`
- For the respective repo CI to run only when the corresponding repo's files change, use Github's Action [`paths field`](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#onpushpull_requestpull_request_targetpathspaths-ignore) to configure a workflow to run based on what file paths are changed

```
// window
winget install -e --id GnuWin32.Make

// linux
sudo apt-get install build-essential

make workspace=services/web_push test
```