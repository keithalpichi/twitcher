package twitch

import "net/url"

// Channeler represents the interface for specific properties of Twitch Channels
type Channeler interface {
	GetChannel(Twitcher, url.Values) (Channel, error)
	GetChannels(Twitcher, url.Values) ([]Channel, error)
	GetFollowers(Twitcher, url.Values) ([]Followers, error)
}

// Channel contains meta information about the Twitch Channel
type Channel struct {
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

// ChannelResponse represents a response from Twitch containing an array of Channels
type ChannelResponse struct {
	Data []Channel `json:"data"`
}

// GetChannel gets a Twitch streamer's channel
func (c Client) GetChannel(twitchClient Twitcher, v url.Values) (ch Channel, err error) {
	return
}

// GetChannels gets an array of Twitch streamer channels
func (c Client) GetChannels(twitchClient Twitcher, v url.Values) (ch []Channel, err error) {
	return
}
