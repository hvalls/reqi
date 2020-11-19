package request

import (
	"reqi/requesttpl"
	"testing"
)

func TestExecuteGet(t *testing.T) {
	tpl := requesttpl.New("name", "description", "https://www.google.com", "get", "")
	client := &MockHTTPClient{false, false}
	r := New(tpl, client)

	r.Execute()

	if !client.DoGetCalled {
		t.Errorf("client.DoGet should be called")
	}
}

func TestExecutePost(t *testing.T) {
	tpl := requesttpl.New("name", "description", "https://www.google.com", "post", "{\"msg\": \"Hello\"}")
	client := &MockHTTPClient{false, false}
	r := New(tpl, client)

	r.Execute()

	if !client.DoPostCalled {
		t.Errorf("client.DoPost should be called")
	}
}

type MockHTTPClient struct {
	DoGetCalled  bool
	DoPostCalled bool
}

func (client *MockHTTPClient) DoGet(url string) (string, error) {
	client.DoGetCalled = true
	return "", nil
}

func (client *MockHTTPClient) DoPost(url, body string) (string, error) {
	client.DoPostCalled = true
	return "", nil
}
