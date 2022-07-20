<img src="images/axolgo-logo-transparent.png" width="50%" />

# axolgo-cloud, the Axolotl Cloud Library in Golang
#### Release
<div align="left">
  <a href="ttps://github.com/tchiunam/axolgo-cloud/releases">
    <img alt="Version" src="https://img.shields.io/github/v/release/tchiunam/axolgo-cloud?sort=semver" />
  </a>
  <a href="https://github.com/tchiunam/axolgo-cloud/releases">
    <img alt="Release Date" src="https://img.shields.io/github/release-date/tchiunam/axolgo-cloud" />
  </a>
  <a href="https://pkg.go.dev/github.com/tchiunam/axolgo-cloud">
    <img alt="PkgGoDev" src="https://pkg.go.dev/badge/github.com/tchiunam/axolgo-cloud" />
  </a>
  <img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/tchiunam/axolgo-cloud" />
  <img alt="Language" src="https://img.shields.io/github/languages/count/tchiunam/axolgo-cloud" />
  <img alt="File Count" src="https://img.shields.io/github/directory-file-count/tchiunam/axolgo-cloud" />
  <img alt="Repository Size" src="https://img.shields.io/github/repo-size/tchiunam/axolgo-cloud.svg?label=Repo%20size" />
</div>

#### Code Quality
<div align="left">
  <a href="https://github.com/tchiunam/axolgo-cloud/actions/workflows/go.yml">
    <img alt="Go" src="https://github.com/tchiunam/axolgo-cloud/actions/workflows/go.yml/badge.svg" />
  </a>
  <a href="https://codecov.io/gh/tchiunam/axolgo-cloud">
    <img alt="codecov" src="https://codecov.io/gh/tchiunam/axolgo-cloud/branch/main/graph/badge.svg?token=7Q6I4OXAS8" />
  </a>
  <a href="https://github.com/tchiunam/axolgo-cloud/actions/workflows/codeql-analysis.yml">
    <img alt="CodeQL" src="https://github.com/tchiunam/axolgo-cloud/actions/workflows/codeql-analysis.yml/badge.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/tchiunam/axolgo-cloud">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/tchiunam/axolgo-cloud" />
  </a>
</div>

#### Activity
<div align="left">
  <a href="https://github.com/tchiunam/axolgo-cloud/commits/main">
    <img alt="Last Commit" src="https://img.shields.io/github/last-commit/tchiunam/axolgo-cloud" />
  </a>
  <a href="https://github.com/tchiunam/axolgo-cloud/issues?q=is%3Aissue+is%3Aclosed">
    <img alt="Closed Issues" src="https://img.shields.io/github/issues-closed/tchiunam/axolgo-cloud" />
  </a>
  <a href="https://github.com/tchiunam/axolgo-cloud/pulls?q=is%3Apr+is%3Aclosed">
    <img alt="Closed Pull Requests" src="https://img.shields.io/github/issues-pr-closed/tchiunam/axolgo-cloud" />
  </a>
</div>

#### License
<div align="left">
  <a href="https://opensource.org/licenses/MIT">
    <img alt="License: MIT" src="https://img.shields.io/github/license/tchiunam/axolgo-cloud" />
  </a>
</div>

#### Popularity
<div align="left">
  <a href="https://sourcegraph.com/github.com/tchiunam/axolgo-cloud?badge">
    <img alt="Sourcegraph" src="https://sourcegraph.com/github.com/tchiunam/axolgo-cloud/-/badge.svg" />
  </a>
  <img alt="Repo Stars" src="https://img.shields.io/github/stars/tchiunam/axolgo-cloud?style=social" />
  <img alt="Watchers" src="https://img.shields.io/github/watchers/tchiunam/axolgo-cloud?style=social" />
</div>

<br />
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

## Run test
To run test:
```
go test ./...
```

To run test with coverage result:
```
go test -coverpkg=./... ./...
```

## Test report
## Code Coverage graph
![Code Coverage graph](https://codecov.io/gh/tchiunam/axolgo-cloud/branch/main/graphs/tree.svg?token=7Q6I4OXAS8)

---
#### See more  
1. [axolgo-lib](https://github.com/tchiunam/axolgo-lib) for the base library
2. [axolgo-cli](https://github.com/tchiunam/axolgo-cli) for using Axolgo in command line
