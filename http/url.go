package http

import "net/url"

func ValidateURL(urlStr string) error {
	_, err := url.ParseRequestURI(urlStr)
	return err
}
