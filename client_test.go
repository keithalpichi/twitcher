package twitch

import (
	"net/url"
	"testing"
	"time"
)

type mockClient struct {
	Client
}

func (mC mockClient) Request(uri string, v url.Values) (resp []byte, err error) {
	return
}

func mockNewClient(valid bool) (c mockClient) {
	if !valid {
		return
	}
	c.Token.Date = time.Now()
	c.Secret = "secret"
	c.ID = "123abc"
	c.Token.AccessToken = "123"
	c.Token.RefreshToken = "456"
	c.Token.Exp = int64(123)
	c.Token.Scope = make([]string, 0)
	return
}

func TestRequestWithInvalidFields(t *testing.T) {
	testTable := []Client{
		Client{},
		Client{Secret: "secret", ID: "123abc"},
		Client{
			Secret: "secret",
			ID:     "123abc",
			Token: Token{
				Date:         time.Now(),
				AccessToken:  "123",
				RefreshToken: "456",
				Exp:          int64(123),
				Scope:        make([]string, 0),
			},
		},
	}
	v := url.Values{}
	for i, testClient := range testTable {
		if _, err := testClient.Request("nowhere", v); err == nil {
			t.Errorf("twitch.Client.Request should return error due to invalid twitch.Client fields. Error in index %d of testTable", i)
		}
	}
}

func TestRequestWithValidFields(t *testing.T) {
	c := mockNewClient(true)
	v := url.Values{}
	if _, err := c.Request("nowhere", v); err == ErrInvalidClient {
		t.Errorf("twitch.Client.Reqeust should return no error when fields are valid. Got %s", err.Error())
	}
}
