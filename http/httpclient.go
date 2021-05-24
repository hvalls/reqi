package http

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const contentType = "application/json; charset=utf-8"

type HTTPClient interface {
	DoGet(url string, headers []*HTTPHeader) (string, error)
	DoPost(url, body string, headers []*HTTPHeader) (string, error)
	DoPut(url, body string, headers []*HTTPHeader) (string, error)
}

type DefaultHTTPClient struct{}

func (client *DefaultHTTPClient) DoGet(url string, headers []*HTTPHeader) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	return prepareAndExec(req, headers)
}

func (client *DefaultHTTPClient) DoPost(url, body string, headers []*HTTPHeader) (string, error) {

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	return prepareAndExec(req, headers)
}

func (client *DefaultHTTPClient) DoPut(url, body string, headers []*HTTPHeader) (string, error) {
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	return prepareAndExec(req, headers)
}

func prepareAndExec(req *http.Request, headers []*HTTPHeader) (string, error) {
	for _, h := range headers {
		req.Header.Set(h.Name, h.Value)
	}
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
