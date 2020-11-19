package httpclient

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type HTTPClient interface {
	DoGet(url string) (string, error)
	DoPost(url, body string) (string, error)
}

type DefaultHTTPClient struct{}

func (client *DefaultHTTPClient) DoGet(url string) (string, error) {
	return getBody(http.Get(url))
}

func (client *DefaultHTTPClient) DoPost(url, body string) (string, error) {
	return getBody(http.Post(url, "application/json", strings.NewReader(body)))
}

func New() HTTPClient {
	return &DefaultHTTPClient{}
}

func getBody(response *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
