package testhttp

import (
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c Client) RequestHello(name string) (string, error) {
	targetUrl, _ := url.JoinPath(c.url, "/hello/", name)
	res, err := http.Get(targetUrl)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
