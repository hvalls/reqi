package http

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const contentType = "application/json; charset=utf-8"

type HTTPClient interface {
	DoGet(url string) (string, error)
	DoPost(url, body string) (string, error)
	DoPut(url, body string) (string, error)
}

type DefaultHTTPClient struct{}

func (client *DefaultHTTPClient) DoGet(url string) (string, error) {
	return getBody(http.Get(url))
}

func (client *DefaultHTTPClient) DoPost(url, body string) (string, error) {
	return getBody(http.Post(url, contentType, strings.NewReader(body)))
}

func (client *DefaultHTTPClient) DoPut(url, body string) (string, error) {
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", contentType)
	cli := &http.Client{}
	return getBody(cli.Do(req))
}

func NewClient() HTTPClient {
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
