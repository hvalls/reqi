package requesttpl

import (
	"reqi/http"

	"gopkg.in/yaml.v2"
)

type RequestTpl struct {
	Name        string
	Description string
	URL         string
	Method      http.HTTPMethod
	Body        string
	Headers     []*http.HTTPHeader
}

func New(name string, description string, url string, method http.HTTPMethod, body string, headers []*http.HTTPHeader) *RequestTpl {
	return &RequestTpl{name, description, url, method, body, headers}
}

func NewYaml(yml []byte) (*RequestTpl, error) {
	var r *RequestTpl
	err := yaml.Unmarshal(yml, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (tpl *RequestTpl) String() (string, error) {
	content, err := yaml.Marshal(tpl)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
