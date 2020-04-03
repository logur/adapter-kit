# Logur adapter for [Go Kit logger](https://github.com/go-kit/kit/tree/master/log)

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/logur/adapter-kit/CI?style=flat-square)](https://github.com/logur/adapter-kit/actions?query=workflow%3ACI)
[![Codecov](https://img.shields.io/codecov/c/github/logur/adapter-kit?style=flat-square)](https://codecov.io/gh/logur/adapter-kit)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/adapter/kit?style=flat-square)](https://goreportcard.com/report/logur.dev/adapter/kit)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/logur.dev/adapter/kit)


## Installation

```bash
go get logur.dev/adapter/kit
```


## Usage

```go
package main

import (
	"github.com/go-kit/kit/log"
	kitadapter "logur.dev/adapter/kit"
)

func main() {
	logger := kitadapter.New(log.NewNopLogger())
}
```


## Development

When all coding and testing is done, please run the test suite:

```bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
