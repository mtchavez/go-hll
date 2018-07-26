# go-hll

[![Build Status](https://travis-ci.org/mtchavez/go-hll.svg?branch=master)](https://travis-ci.org/mtchavez/go-hll)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mtchavez/go-hll)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtchavez/go-hll)](https://goreportcard.com/report/github.com/mtchavez/go-hll)
[![Maintainability](https://api.codeclimate.com/v1/badges/cbdf4f5b5cfa83ad2030/maintainability)](https://codeclimate.com/github/mtchavez/go-hll/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/cbdf4f5b5cfa83ad2030/test_coverage)](https://codeclimate.com/github/mtchavez/go-hll/test_coverage)

Go implementation of Hyper Log Log

## Install

```go
go get github.com/mtchavez/go-hll/hll
```

## Usage

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

## Tests

Run tests using ```go test```.

## License

Written by Chavez

Released under the MIT License: http://www.opensource.org/licenses/mit-license.php
