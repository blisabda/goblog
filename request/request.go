package request

import (
	"io"
	"net/http"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

type (
	WebClientFactory interface {
		Create(timeout time.Duration) Client
	}

	Client interface {
		Get(url string, headers http.Header) (*http.Response, error)
		Post(url string, body io.Reader, headers http.Header) (*http.Response, error)
		Put(url string, body io.Reader, headers http.Header) (*http.Response, error)
		Patch(url string, body io.Reader, headers http.Header) (*http.Response, error)
		Delete(url string, headers http.Header) (*http.Response, error)
		Do(req *http.Request) (*http.Response, error)
	}
)

type (
	clientFactory struct {
		timeout time.Duration
	}

	client struct {
		Doer httpclient.Client
	}
)

func (c *client) Get(url string, headers http.Header) (*http.Response, error) {
	return c.Doer.Get(url, headers)
}
func (c *client) Post(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.Doer.Post(url, body, headers)
}
func (c *client) Put(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.Doer.Put(url, body, headers)
}
func (c *client) Patch(url string, body io.Reader, headers http.Header) (*http.Response, error) {
	return c.Doer.Patch(url, body, headers)
}
func (c *client) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.Doer.Delete(url, headers)
}
func (c *client) Do(req *http.Request) (*http.Response, error) {
	return c.Doer.Do(req)
}

func (cf *clientFactory) Create(timeout time.Duration) Client {
	return &client{Doer: *httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))}
}

func NewClientFactory() WebClientFactory {
	return &clientFactory{}
}
