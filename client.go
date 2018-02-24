package twitch

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/franela/goreq"
)

const (
	twitchAPI = "https://api.twitch.tv/"
)

var (
	// ErrNotFound is an error for a resource that isn't found
	ErrNotFound = errors.New("not found")

	// ErrInvalidReq is an error for a invalid request
	ErrInvalidReq = errors.New("invalid request")

	// ErrReqTimeout is an error for a request timeout
	ErrReqTimeout = errors.New("request timeout")

	// ErrTwitch is an error for a server error returned by Twitch
	ErrTwitch = errors.New("Twitch server error")

	// ErrAuth is an error for an invalid request due to an expired token or invalid credentials
	ErrAuth = errors.New("invalid token or credentials")

	// ErrInvalidClient is an error for an invalid twitch.Client struct
	ErrInvalidClient = errors.New("invalid twitch.Client")
)

// Twitcher represents the interface for the Twitch client
type Twitcher interface {
	Request(string, url.Values) ([]byte, error)
	GetAppAccessToken() (Token, error)
	GetRefreshToken() (Token, error)
}

// Token represents a twitch token object returned from a token reqeust
type Token struct {
	Date         time.Time
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Exp          int64    `json:"expires_in"`
	Scope        []string `json:"scope"`
}

// Client represents the Twitch client
type Client struct {
	Token
	Secret string
	ID     string
}

// NewClient initializes a `Client` struct with `s` as the app secret and `id` as the app ID
func NewClient(s, id string) (c *Client) {
	c.Secret = s
	c.ID = id
	return
}

// Request gets a resource from Twitch at the `uri` resource with `v` query string parameters
// It returns a byte array and an error
func (c Client) Request(uri string, v url.Values) (resp []byte, err error) {
	validToken := func(t Token) (b bool) {
		if (t.Date == time.Time{} || t.AccessToken == "" || t.RefreshToken == "" || t.Exp == int64(0)) {
			return
		}
		return true
	}
	if c.Secret == "" || c.ID == "" || !validToken(c.Token) {
		err = ErrInvalidClient
		return
	}

	res, rErr := goreq.Request{
		Method:      "GET",
		Uri:         twitchAPI + uri,
		QueryString: v,
	}.WithHeader("Authorization", fmt.Sprintf("Bearer %s", c.Token.AccessToken)).Do()
	defer res.Body.Close()
	if rErr != nil {
		err = rErr
		return
	}

	switch res.StatusCode {
	case 200:
		resp, err = ioutil.ReadAll(res.Body)
	case 401:
		err = ErrAuth
	case 404:
		err = ErrNotFound
	case 400, 403:
		err = fmt.Errorf("bad request with status code %d", res.StatusCode)
	case 408:
		err = ErrReqTimeout
	default:
		err = ErrTwitch
	}
	return
}

// GetAppAccessToken gets app access token from Twitch
func (c Client) GetAppAccessToken() (t Token, err error) {
	// set t.Date to time.Now().Unix()
	// c.Request, decode into Token
	return
}

// GetRefreshToken gets a new token from Twitch
func (c Client) GetRefreshToken() (t Token, err error) {
	// c.Request, decode into token
	return
}

// IsExpired returns true if the `t.AccessToken` is expired, false otherwise
func (t Token) IsExpired() (b bool) {
	return t.Date.Unix()+t.Exp > time.Now().Unix()
}

func validToken(t Token) (b bool) {
	if (t.Date == time.Time{} || t.AccessToken == "" || t.RefreshToken == "" || t.Exp == int64(0)) {
		return
	}
	return true
}
