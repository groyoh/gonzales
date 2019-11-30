package gonzales_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/groyoh/gonzales"
)

func TestHeader(t *testing.T) {
	c := qt.New(t)
	g := gonzales.Header("Foo", "Bar")

	s := httptest.NewServer(g)
	resp, err := http.Get(s.URL)
	c.Assert(err, qt.IsNil)

	c.Assert(resp.Header.Get("Foo"), qt.Equals, "Bar")
}

func TestStatus(t *testing.T) {
	c := qt.New(t)
	g := gonzales.Status(http.StatusNotFound)

	s := httptest.NewServer(g)
	resp, err := http.Get(s.URL)
	c.Assert(err, qt.IsNil)

	c.Assert(resp.StatusCode, qt.Equals, http.StatusNotFound)
}

func TestBody(t *testing.T) {
	c := qt.New(t)
	g := gonzales.Body("Hello")

	s := httptest.NewServer(g)
	resp, err := http.Get(s.URL)
	c.Assert(err, qt.IsNil)

	c.Assert(resp.StatusCode, qt.Equals, http.StatusOK)
	c.Assert(mustGetBody(t, resp), qt.Equals, "Hello")
}

func TestMirrorBody(t *testing.T) {
	c := qt.New(t)
	g := gonzales.MirrorBody()

	s := httptest.NewServer(g)
	resp, err := http.Post(s.URL, "application/json", strings.NewReader(`[1,2,3]`))
	c.Assert(err, qt.IsNil)

	c.Assert(resp.StatusCode, qt.Equals, http.StatusOK)
	c.Assert(mustGetBody(t, resp), qt.Equals, `[1,2,3]`)
}

func TestMirrorBody_EmptyBody(t *testing.T) {
	c := qt.New(t)
	g := gonzales.MirrorBody()

	s := httptest.NewServer(g)
	resp, err := http.Get(s.URL)
	c.Assert(err, qt.IsNil)

	c.Assert(resp.StatusCode, qt.Equals, http.StatusOK)
	c.Assert(mustGetBody(t, resp), qt.Equals, ``)
}

func TestGonzales_MirrorBody(t *testing.T) {
	c := qt.New(t)
	g := gonzales.New().MirrorBody()

	s := httptest.NewServer(g)
	resp, err := http.Post(s.URL, "application/json", strings.NewReader(`[1,2,3]`))
	c.Assert(err, qt.IsNil)

	c.Assert(resp.StatusCode, qt.Equals, http.StatusOK)
	c.Assert(mustGetBody(t, resp), qt.Equals, `[1,2,3]`)
}

func TestGonzales_Header(t *testing.T) {
	c := qt.New(t)
	g := gonzales.New()
	g.Header("Foo", "Bar")

	s := httptest.NewServer(g)
	resp, err := http.Get(s.URL)
	c.Assert(err, qt.IsNil)

	c.Assert(resp.Header.Get("Foo"), qt.Equals, "Bar")
}

func TestGonzales_Status(t *testing.T) {
	c := qt.New(t)
	g := gonzales.New()
	g.Status(http.StatusNotFound)

	s := httptest.NewServer(g)
	resp, err := http.Get(s.URL)
	c.Assert(err, qt.IsNil)

	c.Assert(resp.StatusCode, qt.Equals, http.StatusNotFound)
}

func TestGonzales_Body(t *testing.T) {
	c := qt.New(t)
	g := gonzales.New()
	g = g.Body("Hello")

	s := httptest.NewServer(g)
	resp, err := http.Get(s.URL)
	c.Assert(err, qt.IsNil)

	c.Assert(resp.StatusCode, qt.Equals, http.StatusOK)
	c.Assert(mustGetBody(t, resp), qt.Equals, "Hello")
}

func TestGonzales_MirrorAllHeaders(t *testing.T) {
	c := qt.New(t)
	g := gonzales.New()
	g = g.MirrorAllHeaders()

	s := httptest.NewServer(g)

	client := &http.Client{}
	req, err := http.NewRequest("GET", s.URL, nil)
	c.Assert(err, qt.IsNil)

	req.Header.Add("Foo", "Bar")
	req.Header.Add("Bar", "Foo")

	resp, err := client.Do(req)
	c.Assert(err, qt.IsNil)

	c.Assert(resp.Header.Get("Foo"), qt.Equals, "Bar")
	c.Assert(resp.Header.Get("Bar"), qt.Equals, "Foo")
}

func TestGonzales_MirrorHeader(t *testing.T) {
	c := qt.New(t)
	g := gonzales.New()
	g = g.MirrorHeader("Foo", "Bar")

	s := httptest.NewServer(g)

	client := &http.Client{}
	req, err := http.NewRequest("GET", s.URL, nil)
	c.Assert(err, qt.IsNil)
	req.Header.Add("Foo", "Bar")
	req.Header.Add("Bar", "Foo")
	req.Header.Add("FooBar", "FooBar")

	resp, err := client.Do(req)
	c.Assert(err, qt.IsNil)

	c.Assert(resp.Header.Get("Foo"), qt.Equals, "Bar")
	c.Assert(resp.Header.Get("Bar"), qt.Equals, "Foo")
	c.Assert(resp.Header.Get("FooBar"), qt.Equals, "")
}

func TestGonzales_chaining(t *testing.T) {
	c := qt.New(t)
	g := gonzales.New()
	g.Body("Hello").Status(http.StatusNotFound).Header("Foo", "Bar")

	s := httptest.NewServer(g)
	resp, err := http.Get(s.URL)
	c.Assert(err, qt.IsNil)

	c.Assert(resp.StatusCode, qt.Equals, http.StatusNotFound)
	c.Assert(mustGetBody(t, resp), qt.Equals, "Hello")
	c.Assert(resp.Header.Get("Foo"), qt.Equals, "Bar")
}

func mustGetBody(t *testing.T, resp *http.Response) string {
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
