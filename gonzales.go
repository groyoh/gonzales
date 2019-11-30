package gonzales

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Gonzales is an http handler with convinient methods.
type Gonzales struct {
	body             string
	header           http.Header
	status           int
	mirrorAllHeaders bool
	mirrorHeaders    map[string]bool
	mirrorBody       bool
}

// New creates a new Gonzales struct.
func New() *Gonzales {
	return &Gonzales{
		header:           http.Header{},
		mirrorAllHeaders: false,
		mirrorHeaders:    make(map[string]bool),
	}
}

// Header creates a new Gonzales struct while setting a header.
func Header(key, value string) *Gonzales {
	return New().Header(key, value)
}

// Body creates a new Gonzales struct while setting its body.
func Body(body string) *Gonzales {
	return New().Body(body)
}

// MirrorBody creates a new Gonzales struct while configuring the handler to
// mirror the request body into the reponse.
func MirrorBody() *Gonzales {
	return New().MirrorBody()
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

// MirrorAllHeaders sets the handler to mirror all the headers
// from the request in the response.
func (g *Gonzales) MirrorAllHeaders() *Gonzales {
	g.mirrorAllHeaders = true
	return g
}

// MirrorHeader set the handler to mirror specific headers from
// the request in the response.
func (g *Gonzales) MirrorHeader(names ...string) *Gonzales {
	for _, name := range names {
		lowerCaseName := strings.ToLower(name)
		g.mirrorHeaders[lowerCaseName] = true
	}
	return g
}

// MirrorBody set the handler to mirror the body from
// the request in the response.
func (g *Gonzales) MirrorBody() *Gonzales {
	g.mirrorBody = true
	return g
}

// ServerHTTP allows Gonzales to implement the http.Handler interface.
func (g *Gonzales) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseHeader := w.Header()
	if g.mirrorAllHeaders {
		g.copyHeaders(r.Header, responseHeader, allHeaders)
	} else {
		g.copyHeaders(r.Header, responseHeader, func(name string) bool {
			lowerCaseName := strings.ToLower(name)
			if _, ok := g.mirrorHeaders[lowerCaseName]; ok {
				return true
			}
			return false
		})
	}
	g.copyHeaders(g.header, responseHeader, allHeaders)

	if g.status != 0 {
		w.WriteHeader(g.status)
	}

	body := []byte(g.body)
	if g.mirrorBody {
		body, _ = ioutil.ReadAll(r.Body)
	}

	w.Write(body)
}

var allHeaders = func(string) bool {
	return true
}

func (g *Gonzales) copyHeaders(in http.Header, out http.Header, assert func(string) bool) {
	for k, values := range in {
		if assert(k) {
			for _, v := range values {
				out.Add(k, v)
			}
		}
	}
}
