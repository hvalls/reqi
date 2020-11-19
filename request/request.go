package request

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reqi/requesttpl"
	"strings"
)

type Request struct {
	Template *requesttpl.RequestTpl
}

func New(tpl *requesttpl.RequestTpl) *Request {
	return &Request{tpl}
}

func (r *Request) Execute() {
	method := strings.ToLower(r.Template.Method)
	if method == "get" {
		r, err := http.Get(r.Template.URL)
		if err != nil {
			log.Fatal(err)
		} else {
			b, _ := ioutil.ReadAll(r.Body)
			fmt.Println(string(b))
		}
		return
	}
	if method == "post" {
		r, err := http.Post(r.Template.URL, "application/json", strings.NewReader(r.Template.Body))
		if err != nil {
			log.Fatal(err)
		} else {
			b, _ := ioutil.ReadAll(r.Body)
			fmt.Println(string(b))
		}
		return
	}
}
