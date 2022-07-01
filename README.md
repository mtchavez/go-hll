# go-hll

[![Latest Version](http://img.shields.io/github/release/mtchavez/go-hll.svg?style=flat-square)](https://github.com/mtchavez/go-hll/releases)
[![Test](https://github.com/mtchavez/go-hll/actions/workflows/test.yml/badge.svg)](https://github.com/mtchavez/go-hll/actions/workflows/test.yml)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mtchavez/go-hll)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtchavez/go-hll)](https://goreportcard.com/report/github.com/mtchavez/go-hll)
[![Maintainability](https://api.codeclimate.com/v1/badges/cbdf4f5b5cfa83ad2030/maintainability)](https://codeclimate.com/github/mtchavez/go-hll/maintainability)
[![Test Coverage](https://codecov.io/gh/mtchavez/go-hll/branch/master/graph/badge.svg?token=YBShFgosRF)](https://codecov.io/gh/mtchavez/go-hll)

Go implementation of Hyper Log Log

## Install

```go
go get github.com/mtchavez/go-hll/hll
```

## Usage

- [New](#create-new)
- [New With Default Error](#new-with-default-error)
- [Adding](#adding)
- [Count](#count)

### Create new

New hyper log log table with a desired error

```go
package main

import (
  "github.com/mtchavez/go-hll/hll"
)

func main() {
  hll := NewWithErr(0.065)
}
```

### New with default error

```go
package main

import (
  "github.com/mtchavez/go-hll/hll"
)

func main() {
  // Uses DefaultErr of 0.065
  hll := New()
  hll.Add("foo")
}
```

### Adding

Add some words to the table

```go
package main

import (
  "github.com/mtchavez/go-hll/hll"
)

func main() {
  hll := NewWithErr(0.065)
  words := []string{"apple", "this is bananas", "kiwi kiwi kiwi", "Peach is a peach", "apple banana peach wiki pear"}
  for _, word := range words {
    hll.Add(word)
  }
}
```

### Count

Get the count and calculate the error of your hyper log log

```go
hll := NewWithErr(0.065)
  words := []string{"apple", "this is bananas", "kiwi kiwi kiwi", "Peach is a peach", "apple banana peach wiki pear"}
  for _, word := range words {
    hll.Add(word)
  }
  count := hll.Count()
  err := float32((count - uint32(len(words))) / uint32(len(words)) / 100.0)
  fmt.Printf("\nCount: %d\nError: %f%%\n\n", count, err)

```

## Documentation

Docs can be generated locally using ```godoc``` or go to [godoc.org](http://godoc.org/github.com/mtchavez/go-hll/hll)

## Benchmarks

```txt
# Updated: 2022-07-01

goos: darwin
goarch: arm64
pkg: github.com/mtchavez/go-hll/hll

BenchmarkHllNew-8         103207             11357 ns/op
BenchmarkHllAdd-8       15996489             74.95 ns/op
BenchmarkHllCount-8       156948              7598 ns/op
```

## Tests

Run tests using ```go test```.

## License

Written by Chavez

Released under the MIT License: <http://www.opensource.org/licenses/mit-license.php>
