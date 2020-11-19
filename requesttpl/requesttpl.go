package requesttpl

type RequestTpl struct {
	Name        string
	Description string
	URL         string
	Method      string
	Body        string
}

func New(name, description, url, method, body string) *RequestTpl {
	return &RequestTpl{name, description, url, method, body}
}
