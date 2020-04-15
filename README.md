# go-base62

[![test](https://github.com/kenkyu392/go-base62/workflows/test/badge.svg?branch=master)](https://github.com/kenkyu392/go-base62)
[![codecov](https://codecov.io/gh/kenkyu392/go-base62/branch/master/graph/badge.svg)](https://codecov.io/gh/kenkyu392/go-base62)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-00ADD8?logo=go)](https://pkg.go.dev/github.com/kenkyu392/go-base62)
[![go report card](https://goreportcard.com/badge/github.com/kenkyu392/go-base62)](https://goreportcard.com/report/github.com/kenkyu392/go-base62)
[![license](https://img.shields.io/github/license/kenkyu392/go-base62)](LICENSE)

Package base62 implements base62 encoding.

## Installation

```
go get -u github.com/kenkyu392/go-base62
```

## Usage

```go
package main

import (
	"github.com/kenkyu392/go-base62"
)

func main() {
	encoded := base62.StdEncoding.EncodeToString([]byte("Hello world!"))
	println(encoded) // T8dgcjRGuYUueWht
}
```

## License

[MIT](LICENSE)
