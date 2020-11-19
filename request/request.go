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
	if r.Template.Method == "get" {
		return r.client.DoGet(r.Template.URL)
	}
	if r.Template.Method == "post" {
		return r.client.DoPost(r.Template.URL, r.Template.Body)
	}
	return "", errors.New("http method not supported")
}
