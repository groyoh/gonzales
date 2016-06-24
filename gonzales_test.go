package gonzales

import (
	"github.com/nbio/st"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHeader(t *testing.T) {
	g := Header("Foo", "Bar")
	w, req := prepareRequest()

	g.ServeHTTP(w, req)
	st.Expect(t, w.Header().Get("Foo"), "Bar")
}

func TestStatus(t *testing.T) {
	g := Status(http.StatusNotFound)
	w, req := prepareRequest()

	g.ServeHTTP(w, req)
	st.Expect(t, w.Code, http.StatusNotFound)
}

func TestBody(t *testing.T) {
	g := Body("Hello")
	w, req := prepareRequest()

	g.ServeHTTP(w, req)
	st.Expect(t, w.Body.String(), "Hello")
}

func TestGonzales_Header(t *testing.T) {
	g := New()
	returnedValue := g.Header("Foo", "Bar")
	w, req := prepareRequest()

	g.ServeHTTP(w, req)
	st.Expect(t, w.Header().Get("Foo"), "Bar")
	st.Expect(t, returnedValue, g)
}

func TestGonzales_Status(t *testing.T) {
	g := New()
	returnedValue := g.Status(http.StatusNotFound)
	w, req := prepareRequest()

	g.ServeHTTP(w, req)
	st.Expect(t, w.Code, http.StatusNotFound)
	st.Expect(t, returnedValue, g)
}

func TestGonzales_Body(t *testing.T) {
	g := New()
	returnedValue := g.Body("Hello")
	w, req := prepareRequest()

	g.ServeHTTP(w, req)
	st.Expect(t, w.Body.String(), "Hello")
	st.Expect(t, returnedValue, g)
}

func TestGonzales_chaining(t *testing.T) {
	g := New()
	g.Body("Hello").Status(http.StatusNotFound).Header("Foo", "Bar")

	w, req := prepareRequest()

	g.ServeHTTP(w, req)
	st.Expect(t, w.Body.String(), "Hello")
	st.Expect(t, w.Code, http.StatusNotFound)
	st.Expect(t, w.Header().Get("Foo"), "Bar")
}

func prepareRequest() (*httptest.ResponseRecorder, *http.Request) {
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	return w, req
}
