# Logur adapter for [Go Kit logger](https://github.com/go-kit/kit/tree/master/log)

[![CircleCI](https://circleci.com/gh/logur/adapter-kit.svg?style=svg)](https://circleci.com/gh/logur/adapter-kit)
[![Coverage](https://gocover.io/_badge/logur.dev/adapter/kit)](https://gocover.io/logur.dev/adapter/kit)
[![Go Report Card](https://goreportcard.com/badge/logur.dev/adapter/kit?style=flat-square)](https://goreportcard.com/report/logur.dev/adapter/kit)
[![GolangCI](https://golangci.com/badges/github.com/logur/adapter-kit.svg)](https://golangci.com/r/github.com/logur/adapter-kit)
[![Go Version](https://img.shields.io/badge/go%20version-%3E=1.11-61CFDD.svg?style=flat-square)](https://github.com/logur/adapter-kit)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/logur.dev/adapter/kit)


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

``` bash
$ make check
```


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
