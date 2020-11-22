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

func NewYaml(yml []byte) (*RequestTpl, error) {
	var values map[string]string
	err := yaml.Unmarshal(yml, &values)
	if err != nil {
		return nil, err
	}
	return newMap(values)
}

func newMap(values map[string]string) (*RequestTpl, error) {
	method, err := http.Method(values["method"])
	if err != nil {
		return nil, err
	}
	url := values["url"]
	err = http.ValidateURL(url)
	if err != nil {
		return nil, err
	}
	return New(values["name"], values["description"], url, method, values["body"]), nil
}

func (tpl *RequestTpl) String() (string, error) {
	content, err := yaml.Marshal(tpl)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
