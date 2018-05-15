/*
Package twitcher is a pure Go client package for interacting with Twitch's New Twitch API (2018). It includes high-level API's to make simple requests for Twitch resources
and low-level API's to let you handle the requests and responses.
*/
package twitcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// EndpointBase is the base URL of Twitch developer API
	EndpointBase = "https://api.twitch.tv/"

	// EndPointUsers is the URL endpoint for users
	EndPointUsers = "https://api.twitch.tv/helix/users"

	// EndPointFollowers is the URL endpoint for user followers
	EndPointFollowers = "https://api.twitch.tv/helix/users/follows"

	// EndPointVideos is the URL endpoint for videos
	EndPointVideos = "https://api.twitch.tv/helix/videos"

	// EndPointStreams is the URL endpoint for streams
	EndPointStreams = "https://api.twitch.tv/helix/streams"

	// EndPointGames is the URL endpoint for games
	EndPointGames = "https://api.twitch.tv/helix/games"

	// EndPointAppAccessTokens is the URL endpoint for requesting app access tokens
	EndPointAppAccessTokens = "https://api.twitch.tv/kraken/oauth2/token"
)

var (
	// ErrNotFound is an error for a resource that isn't found
	ErrNotFound = errors.New("not found")

	// ErrInvalidReq is an error for a invalid request
	ErrInvalidReq = errors.New("invalid request")

	// ErrReqTimeout is an error for a request timeout
	ErrReqTimeout = errors.New("request timeout")

	// ErrBadRequest is an error for a bad request
	ErrBadRequest = errors.New("request timeout")

	// ErrForbidden is an error for a forbidden request
	ErrForbidden = errors.New("request forbidden")

	// ErrTwitch is an error for a server error returned by Twitch
	ErrTwitch = errors.New("Twitch server error")

	// ErrAuth is an error for an invalid request due to an expired token or invalid credentials
	ErrAuth = errors.New("invalid token or credentials")

	// ErrInvalidClient is an error for an invalid twitch.Client struct
	ErrInvalidClient = errors.New("invalid twitch.Client")
)

// Communicator represents the interface for the Twitch client
type Communicator interface {
	Do(*http.Request) (*http.Response, error)
}

// Client represents the Twitch client
type Client struct {
	HTTP Communicator
	Token
	AppConfig
}

// AppConfig represents the Twitch application secret and ID
type AppConfig struct {
	Secret string
	ID     string
}

// ClientConfig is the configuration struct provided to `twitcher.NewClient` to initialize a new client
type ClientConfig struct {
	HTTP Communicator
	AppConfig
}

// Request contains information about a future HTTP request; such as the http.Request and the destination URL
type Request struct {
	HTTP http.Request
	URL  string
}

func defaultComunicator() Communicator {
	return &http.Client{
		Timeout: time.Second * 30,
	}
}

// NewClient initializes a `Client` struct with the app secret and ID along with a http client
func NewClient(cf ClientConfig) (c *Client) {
	c.Secret = cf.Secret
	c.ID = cf.ID
	if cf.HTTP == nil {
		c.HTTP = defaultComunicator()
	} else {
		c.HTTP = cf.HTTP
	}
	return
}

// Request gets a resource from Twitch with the values from `opts`
// It returns a byte array and an error
func (c Client) Request(opts Request) (resp []byte, err error) {
	if !c.validCredentials() || !c.Token.valid() {
		err = ErrInvalidClient
		return
	}
	if !opts.validRequest() {
		err = ErrInvalidReq
		return
	}

	req, err := http.NewRequest(opts.HTTP.Method, fmt.Sprintf("%s?%s", opts.URL, opts.HTTP.Form.Encode()), nil)
	if err != nil {
		return
	}
	req.Header.Add("Authorization", c.Token.AuthorizationHeader())
	res, err := c.HTTP.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		resp, err = ioutil.ReadAll(res.Body)
	case http.StatusUnauthorized:
		err = ErrAuth
	case http.StatusNotFound:
		err = ErrNotFound
	case http.StatusBadRequest:
		err = ErrBadRequest
	case http.StatusForbidden:
		err = ErrForbidden
	case http.StatusRequestTimeout:
		err = ErrReqTimeout
	default:
		err = ErrTwitch
	}
	return
}

// AppAccessToken gets app access token from Twitch and puts it in the `c` Client.
// It returns an error. A nil error means the request was successful.
func (c Client) AppAccessToken() (err error) {
	v := url.Values{}
	v.Set("client_id", c.ID)
	v.Set("client_secret", c.Secret)
	v.Set("grant_type", "client_credentials")
	opts := Request{
		HTTP: http.Request{
			Method: http.MethodGet,
			Form:   v,
		},
		URL: EndPointAppAccessTokens,
	}
	bytes, e := c.Request(opts)
	if e != nil {
		err = e
		return
	}
	var t Token
	if e := json.Unmarshal(bytes, &t); err != nil {
		err = e
		return
	}
	c.Token = t
	return
}

func (c Client) validCredentials() (b bool) {
	if c.Secret == "" || c.ID == "" {
		return
	}
	return true
}

func (r Request) validRequest() (b bool) {
	if r.HTTP.Method == "" || r.URL == "" || len(r.HTTP.Form) == 0 {
		return
	}
	return true
}
