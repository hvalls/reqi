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

func (r *Request) Execute() (string, error) {
	switch method := r.Template.Method; method {
	case http.Get:
		return r.client.DoGet(r.Template.URL)
	case http.Post:
		return r.client.DoPost(r.Template.URL, r.Template.Body)
	case http.Put:
		return r.client.DoPut(r.Template.URL, r.Template.Body)
	default:
		return "", errors.New("http method not supported")
	}
}
