Download and install Golang
- `rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz`
- `export PATH=$PATH:/usr/local/go/bin`
- `go version`

Initializing a Golang project
`go mod init example/hello`

Running a Golang plugin
`go run .`

Refreshing dependencies
`go mod tidy`

Official Golang courses
https://go.dev/doc/tutorial/

Reading function in different modules locally:
- `go mod edit -replace example.com/greetings=../greetings`
- `go mod tidy`

Link golang CLI: https://go.dev/ref/mod

Create the module:
`go mod init example.com/greetings`

Testing:
`go test -v`

Compiles the packages, along with their dependencies, but it doesn't install the results.:
- `go build` (In the main/executable module)
- `./hello`

Installing the package in the SO:
- Getting the installation path: `go list -f '{{.Target}}'`
- Exporting the previous path in the SO: `export PATH=$PATH:/previous-path`
- Alternative we can set all the Golang binaries SO path: `go env -w GOBIN=/binary-golang-path` 
- Move the binary to the expected binary path: `go install`
- Executing the `hello` or `binary-name` anywhere