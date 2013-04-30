go-hll
======

Go implementation of Hyper Log Log

## Install

```go
go get https://github.com/mtchavez/go-hll
```

## Example

Setup hash table first

```go
package main

import (
  "go_hll"
  "fmt"
)

func main() {
  info := go_hll.Hll(0.065)
  // Returns Struct with hash table
  fmt.Printf(info.Table)
}
```

Add some words to the table

```go
package main

import (
  "go_hll"
  "fmt"
)

func main() {
  info := go_hll.Hll(0.065)
  words := [5]string {"apple", "this is bananas", "kiwi kiwi kiwi", "Peach is a peach", "apple banana peach wiki pear" }
  for i := 0; i < len(words); i++ {
    go_hll.Add(info, words[i])
  }

  for x, y := range info.Table {
      fmt.Print(x)
      fmt.Print(" - ")
      fmt.Println(y)
  }
}
```

Get count and error

```go
package main

import (
  "go_hll"
  "fmt"
)

func main() {
  info := go_hll.Hll(0.065)
  words := [5]string {"apple", "this is bananas", "kiwi kiwi kiwi", "Peach is a peach", "apple banana peach wiki pear" }
  for i := 0; i < len(words); i++ {
    go_hll.Add(info, words[i])
  }

  for x, y := range info.Table {
      fmt.Print(x)
      fmt.Print(" - ")
      fmt.Println(y)
  }
  count := go_hll.Count(info)
  error := float32((count - uint32(len(words))) / uint32(len(words)) / 100.0)
  fmt.Printf("\n\nCount: %d\nError: %f %%\n\n", count, error)
}

```
