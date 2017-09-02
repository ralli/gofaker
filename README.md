# Gofaker

This library is (more or less) a port of Benjamin Curtis (stympy) faker library to generate fake data (which is itself a port of Perl's Data::Faker library).

## Installing

```bash
go get github.com/ralli/gofaker
```

## Basic Usage

```go
package main

import (
  "fmt"
   "github.com/ralli/gofaker"
)

func main() {
  faker := gofaker.NewFaker("en")
  fmt.Println(faker.Name.Name()) // Prints something like "Herbert Miller"
}
```
