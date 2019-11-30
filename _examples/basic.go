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
