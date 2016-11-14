package gonzales

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

const contentTypeHeader = "Content-Type"
const jsonContentType = "application/json"

// Gonzales is an http handler with convinient methods.
type Gonzales struct {
	body             string
	header           http.Header
	status           int
	mirrorAllHeaders bool
	mirrorHeaders    map[string]bool
	contentType      string
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

// ContentType creates a new Gonzales struct while setting a the Content-Type.
func ContentType(contentType string) *Gonzales {
	return New().ContentType(contentType)
}

// Body creates a new Gonzales struct while setting its body.
func Body(body string) *Gonzales {
	return New().Body(body)
}

// JSON creates a new Gonzales struct while setting its body
// from the given interface and sets its Content-Type to "application/json"..
func JSON(data interface{}) *Gonzales {
	return New().JSON(data)
}

// Status creates a new Gonzales struct while setting is http status.
func Status(status int) *Gonzales {
	return New().Status(status)
}

// MirrorAllHeaders creates a new Gonzales struct while setting
// the handler to return all the headers from the request in the response.
func MirrorAllHeaders() *Gonzales {
	return New().MirrorAllHeaders()
}

// MirrorHeader creates a new Gonzales struct while setting
// the handler to return specific headers from the request in the response.
func MirrorHeader(names ...string) *Gonzales {
	return New().MirrorHeader(names...)
}

// Header sets a header of the handler.
func (g *Gonzales) Header(key, value string) *Gonzales {
	g.header.Add(key, value)
	return g
}

// ContentType creates a new Gonzales struct while setting a the Content-Type.
func (g *Gonzales) ContentType(contentType string) *Gonzales {
	g.contentType = contentType
	return g
}

// Body sets the body of the handler.
func (g *Gonzales) Body(body string) *Gonzales {
	g.body = body
	return g
}

// JSON sets the body of the handler from the given interface and
// sets its Content-Type to "application/json".
func (g *Gonzales) JSON(data interface{}) *Gonzales {
	buf := &bytes.Buffer{}

	switch data.(type) {
	case string:
		buf.WriteString(data.(string))
	case []byte:
		buf.Write(data.([]byte))
	default:
		if err := json.NewEncoder(buf).Encode(data); err != nil {
			return g
		}
	}
	g.body = buf.String()
	g.ContentType(jsonContentType)
	return g
}

// Status sets the http status of the handler.
func (g *Gonzales) Status(status int) *Gonzales {
	g.status = status
	return g
}

// MirrorAllHeaders sets the handler to return all the headers
// from the request in the response.
func (g *Gonzales) MirrorAllHeaders() *Gonzales {
	g.mirrorAllHeaders = true
	return g
}

// MirrorHeader set the handler to return specific headers from
// the request in the response.
func (g *Gonzales) MirrorHeader(names ...string) *Gonzales {
	for _, name := range names {
		lowerCaseName := strings.ToLower(name)
		g.mirrorHeaders[lowerCaseName] = true
	}
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
	if g.contentType != "" {
		responseHeader.Set(contentTypeHeader, g.contentType)
	}
	if g.status != 0 {
		w.WriteHeader(g.status)
	}
	if len(g.body) != 0 {
		w.Write([]byte(g.body))
	}
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
