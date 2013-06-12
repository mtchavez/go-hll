go-hll [![Build Status](https://travis-ci.org/mtchavez/go-hll.png)](https://travis-ci.org/mtchavez/go-hll)
======

Go implementation of Hyper Log Log

## Install

```go
go get https://github.com/mtchavez/go-hll/hll
```

## Usage

### Initializing
Initializing a hyper log log table

```go
package main

import (
  "github.com/mtchavez/go-hll/hll"
)

func main() {
  hll := Initialize(0.065)
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
  hll := Initialize(0.065)
  words := []string{"apple", "this is bananas", "kiwi kiwi kiwi", "Peach is a peach", "apple banana peach wiki pear"}
  for _, word := range words {
    hll.Add(word)
  }
}
```

### Count
Get the count and calculate the error of your hyper log log

```go
hll := Initialize(0.065)
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
