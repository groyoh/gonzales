# Gonzales
[![Build Status](https://travis-ci.org/groyoh/gonzales.svg?branch=master)](https://travis-ci.org/groyoh/gonzales?branch=master)
[![GoDoc](https://godoc.org/github.com/groyoh/gonzales?status.svg)](https://godoc.org/github.com/groyoh/gonzales)
[![Coverage Status](https://coveralls.io/repos/github/groyoh/gonzales/badge.svg?branch=master)](https://coveralls.io/github/groyoh/gonzales?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/groyoh/gonzales)](https://goreportcard.com/report/github.com/groyoh/gonzales)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/groyoh/gonzales/blob/master/LICENSE.md)

This library is meant to provide convinient methods to quickly build http handlers.

## Installation

```
go get -u gopkg.in/groyoh/gonzales.v0
```

## API

See [godoc reference](https://godoc.org/github.com/groyoh/gonzales) for detailed API documentation.

## Examples

```go
package main

import (
  "net/http"

  "gopkg.in/groyoh/gonzales.v0"
)

func main() {
  g := gonzales.Header("Foo", "Bar").
    Status(404).
    Body("Not found").
    MirrorHeader("Foo", "Bar")
  http.Handle("/", g)
  http.ListenAndServe(":8000", nil)
}

```

## License

[MIT](LICENSE.md)

## Development

Clone this repository:
```bash
git clone https://github.com/groyoh/gonzales.git && cd gonzales
```

Install dependencies:
```bash
go get -u -t ./...
```

Run tests:
```bash
go test ./...
```

Lint code:
```bash
golint ./...
```
