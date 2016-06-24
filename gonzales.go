package gonzales

import (
	"net/http"
)

// Gonzales is an http handler with convinient methods.
type Gonzales struct {
	body   string
	header http.Header
	status int
}

// New creates a new Gonzales struct.
func New() *Gonzales {
	return &Gonzales{header: http.Header{}}
}

// Header creates a new Gonzales struct while setting a header.
func Header(key, value string) *Gonzales {
	return New().Header(key, value)
}

// Body creates a new Gonzales struct while setting its body.
func Body(body string) *Gonzales {
	return New().Body(body)
}

// Status creates a new Gonzales struct while setting is http status.
func Status(status int) *Gonzales {
	return New().Status(status)
}

// Header sets a header of the handler.
func (g *Gonzales) Header(key, value string) *Gonzales {
	g.header.Add(key, value)
	return g
}

// Body sets the body of the handler.
func (g *Gonzales) Body(body string) *Gonzales {
	g.body = body
	return g
}

// Status sets the http status of the handler.
func (g *Gonzales) Status(status int) *Gonzales {
	g.status = status
	return g
}

// ServerHTTP allows Gonzales to implement the http.Handler interface.
func (g *Gonzales) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for k, values := range g.header {
		for _, v := range values {
			w.Header().Add(k, v)
		}
	}
	if g.status != 0 {
		w.WriteHeader(g.status)
	}
	if len(g.body) != 0 {
		w.Write([]byte(g.body))
	}
}
