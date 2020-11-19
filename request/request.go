package request

import (
	"errors"
	"reqi/http"
	"reqi/requesttpl"
	"strings"
)

type Request struct {
	Template *requesttpl.RequestTpl
	client   http.HTTPClient
}

func New(tpl *requesttpl.RequestTpl, client http.HTTPClient) *Request {
	return &Request{tpl, client}
}

func (r *Request) Execute() (string, error) {
	method := strings.ToLower(r.Template.Method)
	if method == "get" {
		return r.client.DoGet(r.Template.URL)
	}
	if method == "post" {
		return r.client.DoPost(r.Template.URL, r.Template.Body)
	}
	return "", errors.New("http method not supported")
}
