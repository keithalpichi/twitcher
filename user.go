package twitcher

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// User contains meta information about the Twitch User
type User struct {
	ID              string `json:"id"`
	LoginName       string `json:"login"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	DisplayName     string `json:"display_name"`
	ProfileImageURL string `json:"profile_image_url"`
	ViewCount       int64  `json:"view_count"`
}

// Followers contains meta information about follow relationships between two Twitch streamers
type Followers struct {
	Total int64 `json:"total"`
	Data  []struct {
		From       string `json:"from_id"`
		To         string `json:"to_id"`
		FollowedAt string `json:"followed_at"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// userResponse represents a response from Twitch containing an array of Users
type userResponse struct {
	Data []User `json:"data"`
}

// UserByID gets a Twitch streamer's channel by their user ID
func (c Client) UserByID(id string) (u User, err error) {
	v := url.Values{}
	v.Set("id", id)
	opts := Request{
		HTTP: http.Request{
			Method: http.MethodGet,
			Form:   v,
		},
		URL: EndPointUsers,
	}
	resp, err := c.Request(opts)
	err = json.Unmarshal(resp, &u)
	return
}

// UserByLogin gets a Twitch streamer's channel by their login name
func (c Client) UserByLogin(l string) (u User, err error) {
	return
}
