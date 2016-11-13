package main

import (
  "net/http"

  "gopkg.in/groyoh/gonzales.v0"
)

func main() {
  g := gonzales.Header("Foo", "Bar").
    Status(404).
    Body("Not found").
    MirrorHeader("FooBar", "FooFoo")
  http.Handle("/", g)
  http.ListenAndServe(":8000", nil)
}
