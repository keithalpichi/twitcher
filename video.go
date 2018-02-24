package twitch

import (
	"net/url"
	"time"
)

// Videor represents the interface for specific properties of Twitch Videos
type Videor interface {
	GetVideo(Twitcher, url.Values) (Video, error)
	GetVideos(Twitcher, url.Values) ([]Video, error)
}

// Video contains meta information about a single Twitch video
type Video struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	PublishedAt  time.Time `json:"published_at"`
	ThumbnailURL string    `json:"thumbnail_url"`
	Language     string    `json:"language"`
	ViewCount    int64     `json:"view_count"`
	Duration     string    `json:"duration"` // in format => "6h4m26s"
}

// VideosResponse represents a response from Twitch containing an array of Videos
type VideosResponse struct {
	Data       []Video `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// GetVideo gets a single Twitch video
func (c Client) GetVideo(twitchClient Twitcher, v url.Values) (vid Video, err error) {
	return
}

// GetVideos gets an array of Twitch videos
func (c Client) GetVideos(twitchClient Twitcher, v url.Values) (vid []Video, err error) {
	return
}
