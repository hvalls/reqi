package request

import (
	"reqi/http"
	"reqi/requesttpl"
	"testing"
)

func TestExecuteGet(t *testing.T) {
	tpl := requesttpl.New("name", "description", "https://www.google.com", "get", "", []*http.HTTPHeader{})
	client := &MockHTTPClient{false, false, false}
	r := New(tpl, client)

	r.Execute()

	if !client.DoGetCalled {
		t.Errorf("client.DoGet should be called")
	}
}

func TestExecutePost(t *testing.T) {
	tpl := requesttpl.New("name", "description", "https://www.google.com", "post", "{\"msg\": \"Hello\"}", []*http.HTTPHeader{})
	client := &MockHTTPClient{false, false, false}
	r := New(tpl, client)

	r.Execute()

	if !client.DoPostCalled {
		t.Errorf("client.DoPost should be called")
	}
}

func TestExecutePut(t *testing.T) {
	tpl := requesttpl.New("name", "description", "https://www.google.com", "put", "{\"msg\": \"Hello\"}", []*http.HTTPHeader{})
	client := &MockHTTPClient{false, false, false}
	r := New(tpl, client)

	r.Execute()

	if !client.DoPutCalled {
		t.Errorf("client.DoPut should be called")
	}
}

type MockHTTPClient struct {
	DoGetCalled  bool
	DoPostCalled bool
	DoPutCalled  bool
}

func (client *MockHTTPClient) DoGet(url string, headers []*http.HTTPHeader) (string, error) {
	client.DoGetCalled = true
	return "", nil
}

func (client *MockHTTPClient) DoPost(url, body string, headers []*http.HTTPHeader) (string, error) {
	client.DoPostCalled = true
	return "", nil
}

func (client *MockHTTPClient) DoPut(url, body string, headers []*http.HTTPHeader) (string, error) {
	client.DoPutCalled = true
	return "", nil
}
