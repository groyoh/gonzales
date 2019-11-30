# Gonzales
[![Build Status](https://travis-ci.org/groyoh/gonzales.svg?branch=master)](https://travis-ci.org/groyoh/gonzales?branch=master)
[![GoDoc](https://godoc.org/github.com/groyoh/gonzales?status.svg)](https://godoc.org/github.com/groyoh/gonzales)
[![Coverage Status](https://coveralls.io/repos/github/groyoh/gonzales/badge.svg?branch=master)](https://coveralls.io/github/groyoh/gonzales?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/groyoh/gonzales)](https://goreportcard.com/report/github.com/groyoh/gonzales)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/groyoh/gonzales/blob/master/LICENSE.md)

This library is meant to provide convinient methods to quickly build http handlers.

## Installation

```
go get github.com/groyoh/gonzales
```

## API

See [godoc reference](https://godoc.org/github.com/groyoh/gonzales) for detailed API documentation.

## Examples

Gonzales can be useful when mocking APIs in test using the `httptest` package:

```go
package gonzales_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/groyoh/gonzales"

	qt "github.com/frankban/quicktest"
)

type Repository struct {
	Slug string `json:"slug"`
}

type GithubClient struct {
	http.Client
	BaseURL string
}

func (c *GithubClient) Repositories() ([]Repository, error) {
	var repos []Repository

	resp, err := http.Get(c.BaseURL + "/repositories")
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(&repos)
	return repos, err
}

func TestGithubClient_Repositories(t *testing.T) {
	c := qt.New(t)
	g := gonzales.Body(`[{"slug":"gonzales"}]`)

	s := httptest.NewServer(g)
	client := GithubClient{BaseURL: s.URL}
	repos, err := client.Repositories()

	c.Assert(err, qt.IsNil)
	c.Assert(len(repos), qt.Equals, 1)
	c.Assert(repos[0].Slug, qt.Equals, "gonzales")
}
```

You may also use Gonzales to build static handlers:

```go
package main

import (
  "net/http"

  "github.com/groyoh/gonzales"
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
