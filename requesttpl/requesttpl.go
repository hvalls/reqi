package requesttpl

import (
	"errors"
	"reqi/http"
	"strings"

	"gopkg.in/yaml.v2"
)

type RequestTpl struct {
	Name        string             `yaml:"name"`
	Description string             `yaml:"description"`
	URL         string             `yaml:"url"`
	Method      http.HTTPMethod    `yaml:"method"`
	Body        string             `yaml:"body"`
	Headers     []*http.HTTPHeader `yaml:"headers"`
}

func (rtpl *RequestTpl) ResolveURL(params map[string]string) (string, error) {
	return resolve(rtpl.URL, "url", params)
}

func (rtpl *RequestTpl) ResolveBody(params map[string]string) (string, error) {
	return resolve(rtpl.URL, "body", params)
}

func (rtpl *RequestTpl) ResolveHeaders(params map[string]string) ([]*http.HTTPHeader, error) {
	headers := []*http.HTTPHeader{}
	for _, v := range rtpl.Headers {
		resolved, err := resolve(v.Value, "headers", params)
		if err != nil {
			return nil, err
		}
		headers = append(headers, &http.HTTPHeader{Name: v.Name, Value: resolved})
	}
	return headers, nil
}

func New(name string, description string, url string, method http.HTTPMethod, body string, headers []*http.HTTPHeader) *RequestTpl {
	return &RequestTpl{name, description, url, method, body, headers}
}

func NewYaml(yml string) (*RequestTpl, error) {
	var r *RequestTpl
	err := yaml.Unmarshal([]byte(yml), &r)
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

func resolve(orig string, field string, params map[string]string) (string, error) {
	resolved := orig
	for k, v := range params {
		resolved = strings.ReplaceAll(resolved, "{{"+k+"}}", v)
		resolved = strings.ReplaceAll(resolved, "{{ "+k+" }}", v)
	}
	if strings.Contains(resolved, "{{") {
		return "", errors.New("unresolved parameter in " + field)
	}
	return resolved, nil
}
