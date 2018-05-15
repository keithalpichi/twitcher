package twitcher

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"
)

type mockHTTP struct {
	fakeStatusCode int
	fakeBody       []byte
}

func (mh mockHTTP) Do(*http.Request) (resp *http.Response, err error) {
	resp = &http.Response{
		StatusCode: mh.fakeStatusCode,
		Body:       ioutil.NopCloser(bytes.NewReader(mh.fakeBody)),
	}
	return
}

func validCredsIf(b bool) (ac AppConfig) {
	if b {
		ac.Secret = "valid secret"
		ac.ID = "valid id"
	}
	return
}

func validTokenIf(b bool) (t Token) {
	if b {
		t.Date = time.Unix(testTime, 0)
		t.AccessToken = "valid access token"
		t.Exp = testTime
	}
	return
}

func invalidRequestIf(b bool) (r Request) {
	if !b {
		r.HTTP.Method = http.MethodGet
		r.HTTP.Form = url.Values{"test": []string{"test"}}
		r.URL = EndpointBase
	}
	return
}

func TestValidCredentials(t *testing.T) {
	cases := []struct {
		client   Client
		expected bool
		msg      string
	}{
		{
			expected: false,
			msg:      "invalid app secret",
		},
		{
			client: Client{
				AppConfig: AppConfig{
					Secret: "valid-app-secret",
				},
			},
			expected: false,
			msg:      "invalid app ID",
		},
	}
	for _, c := range cases {
		if got := c.client.validCredentials(); got != c.expected {
			t.Errorf("%s. Got %v, expected %v", c.msg, got, c.expected)
		}
	}
}

func TestValidRequest(t *testing.T) {
	cases := []struct {
		request  Request
		expected bool
		msg      string
	}{
		{
			expected: false,
			msg:      "invalid method",
		},
		{
			request: Request{
				HTTP: http.Request{
					Method: http.MethodGet,
					Form:   url.Values{"test": []string{"test"}},
				},
			},
			expected: false,
			msg:      "invalid URL",
		},
		{
			request: Request{
				HTTP: http.Request{
					Method: http.MethodGet,
				},
				URL: EndpointBase,
			},
			expected: false,
			msg:      "invalid url.values",
		},
	}
	for _, c := range cases {
		if got := c.request.validRequest(); got != c.expected {
			t.Errorf("%s. Got %v, expected %v", c.msg, got, c.expected)
		}
	}
}

func TestInvalidClientRequests(t *testing.T) {
	cases := []struct {
		validClient    bool
		invalidRequest bool
		statusCode     int
		request        Request
		expectedErr    error
		msg            string
	}{
		{
			validClient: false,
			expectedErr: ErrInvalidClient,
			msg:         "invalid credentials",
		},
		{
			validClient:    true,
			invalidRequest: true,
			expectedErr:    ErrInvalidReq,
			msg:            "invalid request method",
		},
		{
			validClient: true,
			statusCode:  http.StatusBadRequest,
			expectedErr: ErrBadRequest,
			msg:         "status code 400",
		},
		{
			validClient: true,
			statusCode:  http.StatusUnauthorized,
			expectedErr: ErrAuth,
			msg:         "status code 401",
		},
		{
			validClient: true,
			statusCode:  http.StatusForbidden,
			expectedErr: ErrForbidden,
			msg:         "status code 403",
		},
		{
			validClient: true,
			statusCode:  http.StatusNotFound,
			expectedErr: ErrNotFound,
			msg:         "status code 404",
		},
		{
			validClient: true,
			statusCode:  http.StatusRequestTimeout,
			expectedErr: ErrReqTimeout,
			msg:         "status code 408",
		},
		{
			validClient: true,
			statusCode:  http.StatusInternalServerError,
			expectedErr: ErrTwitch,
			msg:         "status code 500+",
		},
	}
	for _, c := range cases {
		client := Client{
			HTTP: mockHTTP{
				fakeStatusCode: c.statusCode,
			},
			Token:     validTokenIf(c.validClient),
			AppConfig: validCredsIf(c.validClient),
		}

		req := invalidRequestIf(c.invalidRequest)

		if _, err := client.Request(req); err != c.expectedErr {
			t.Errorf("%s. Got %v, expected %v", c.msg, err, c.expectedErr)
		}
	}
}

func TestValidClientRequests(t *testing.T) {
	cases := []struct {
		request      Request
		expectedResp []byte
		msg          string
	}{
		{
			expectedResp: []byte("valid response"),
			msg:          "valid request",
		},
	}
	for _, c := range cases {
		client := Client{
			HTTP: mockHTTP{
				fakeStatusCode: http.StatusOK,
				fakeBody:       c.expectedResp,
			},
			Token:     validTokenIf(true),
			AppConfig: validCredsIf(true),
		}

		req := invalidRequestIf(false)

		if resp, err := client.Request(req); err != nil {
			t.Errorf("%s. Got %s, expected %s", c.msg, err.Error(), string(c.expectedResp))
		} else if string(resp) != string(c.expectedResp) {
			t.Errorf("%s. Got %s, expected %s", c.msg, string(resp), string(c.expectedResp))
		}
	}
}
