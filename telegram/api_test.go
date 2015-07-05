package telegram

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.Path)
		fmt.Fprint(w, r.URL.Query().Encode())
	}))
	defer ts.Close()

	config := NewConfig("abcxyz", func(c *Config) {
		c.WithUrl(ts.URL)
	})

	api, _ := NewApi(config)
	q := make(map[string]string)
	q["foo"] = "bar"
	q["foo2"] = "bar2"
	result, err := api.get("method", q)

	assert.Equal(t, nil, err, "should have no errors")
	assert.Equal(t, "/abcxyz/method\nfoo=bar&foo2=bar2", string(result), "should pass corret params")
}

func TestPostRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintln(w, r.URL.Path)
		fmt.Fprint(w, r.PostForm.Encode())
	}))
	defer ts.Close()

	config := NewConfig("abcxyz", func(c *Config) {
		c.WithUrl(ts.URL)
	})
	api, _ := NewApi(config)
	q := make(map[string]string)
	q["foo"] = "bar"
	q["foo2"] = "bar2"
	result, err := api.post("method", q)

	assert.Equal(t, nil, err, "should have no errors")
	assert.Equal(t, "/abcxyz/method\nfoo=bar&foo2=bar2", string(result), "should pass corret params")
}
