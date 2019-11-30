package gonzales_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/groyoh/gonzales"
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

// ExampleHTTPTest demonstrates the usage of gonzales when testing http clients.
func Example_httpTest() {
	g := gonzales.Body(`[{"slug":"gonzales"}]`)

	s := httptest.NewServer(g)
	c := GithubClient{BaseURL: s.URL}
	repos, _ := c.Repositories()

	fmt.Println(repos[0].Slug)
	// Output: gonzales
}
