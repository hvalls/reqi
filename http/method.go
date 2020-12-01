package http

import "errors"

type HTTPMethod string

const (
	Get  HTTPMethod = "get"
	Post            = "post"
	Put             = "put"
)

func Method(value string) (HTTPMethod, error) {
	if value == "get" {
		return Get, nil
	}
	if value == "post" {
		return Post, nil
	}
	if value == "put" {
		return Put, nil
	}
	return "", errors.New("method not supported")
}
