# axolgo-cloud, the Axolotl Cloud Library in Golang
This is the cloud library of the Axolotl series in Golang. It is a middle layer for the application or CLI to use AWS and GCP. The basic calls like loading AWS configuration is handled for you. The difficult part of using Google API in Golang is handled for you. The client calls are also wrapped so you can focus on building your business logic. The interfaces and functions are designed to be more friendly to command line use cases.

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

---
#### See more  
1. [axolgo-lib](https://github.com/tchiunam/axolgo-lib) for the base library
2. [axolgo-cli](https://github.com/tchiunam/axolgo-cli) for using Axolgo in command line