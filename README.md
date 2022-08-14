# GO monorepo
Monorepo in golang using [go workspaces](https://go.dev/doc/tutorial/workspaces)

## `libs` vs `services`
A library is shared code that you compile into your application. A service is a shared capability that you access from your application (e.g. APIs) (https://blogs.gartner.com/eric-knipp/2013/03/20/libraries-vs-services/)

## List of projects
| Service       | Description                                                                                          |
|---------------|------------------------------------------------------------------------------------------------------|
| web-push-SaaS | Web Push SaaS allows you to easily send web push notifications to users with just a single API call. |
|               |                                                                                                      |

## Development guide
- When creating a new module, remember to update the `go.work` file with the directory for the sym link to work
- Run the services with `make workspace=<workspace> <cmd>`
```
// window
winget install -e --id GnuWin32.Make

// linux
sudo apt-get install build-essential

make workspace=services/web_push test
```