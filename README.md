<img src="images/axolgo-logo-transparent.png" width="50%" />

# axolgo-cloud, the Axolotl Cloud Library in Golang
[![Go](https://github.com/tchiunam/axolgo-cloud/actions/workflows/go.yml/badge.svg)](https://github.com/tchiunam/axolgo-cloud/actions/workflows/go.yml)
[![Version](https://img.shields.io/github/v/release/tchiunam/axolgo-cloud?sort=semver)](https://github.com/tchiunam/axolgo-cloud/releases)
[![codecov](https://codecov.io/gh/tchiunam/axolgo-cloud/branch/main/graph/badge.svg?token=7Q6I4OXAS8)](https://codecov.io/gh/tchiunam/axolgo-cloud)
[![CodeQL](https://github.com/tchiunam/axolgo-cloud/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/tchiunam/axolgo-cloud/actions/workflows/codeql-analysis.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/tchiunam/axolgo-cloud)](https://goreportcard.com/report/github.com/tchiunam/axolgo-cloud)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tchiunam/axolgo-cloud)](https://pkg.go.dev/github.com/tchiunam/axolgo-cloud)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This is the cloud library of the Axolotl series in Golang. It is a middle layer for the application or CLI to use AWS and GCP. The basic calls like loading AWS configuration is handled for you. The difficult part of using Google API in Golang is handled for you. The client calls are also wrapped so you can focus on building your business logic. The interfaces and functions are designed to be more friendly to command line use cases.

Go package: https://pkg.go.dev/github.com/tchiunam/axolgo-cloud

## Use it with your Go module
To add as dependency for your package or upgrade to the latest version:
```
go get github.com/tchiunam/axolgo-cloud
```

To upgrade or downgrade to a specific version:
```
go get github.com/tchiunam/axolgo-cloud@v1.2.3
```

To remove dependency on your module and downgrade modules:
```
go get github.com/tchiunam/axolgo-cloud@none
```

See 'go help get' or https://golang.org/ref/mod#go-get for details.

## Test report
## Code Coverage graph
![Code Coverage graph](https://codecov.io/gh/tchiunam/axolgo-cloud/branch/main/graphs/tree.svg?token=7Q6I4OXAS8)

---
#### See more  
1. [axolgo-lib](https://github.com/tchiunam/axolgo-lib) for the base library
2. [axolgo-cli](https://github.com/tchiunam/axolgo-cli) for using Axolgo in command line
