package request

import (
	"errors"
	"reqi/http"
	"reqi/requesttpl"
)

type Request struct {
	Template *requesttpl.RequestTpl
	client   http.HTTPClient
}

func New(tpl *requesttpl.RequestTpl, client http.HTTPClient) *Request {
	return &Request{tpl, client}
}

func (r *Request) Execute(params map[string]string) (string, error) {
	url, err := r.Template.ResolveURL(params)
	if err != nil {
		return "", err
	}
	body, err := r.Template.ResolveBody(params)
	if err != nil {
		return "", err
	}
	headers, err := r.Template.ResolveHeaders(params)
	if err != nil {
		return "", err
	}
	switch method := r.Template.Method; method {
	case http.Get:
		return r.client.DoGet(url, headers)
	case http.Post:
		return r.client.DoPost(url, body, headers)
	case http.Put:
		return r.client.DoPut(url, body, headers)
	default:
		return "", errors.New("http method not supported")
	}
}
