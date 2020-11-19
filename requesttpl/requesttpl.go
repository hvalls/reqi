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
}

func New(name string, description string, url string, method http.HTTPMethod, body string) *RequestTpl {
	return &RequestTpl{name, description, url, method, body}
}

func (tpl *RequestTpl) String() (string, error) {
	content, err := yaml.Marshal(tpl)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
