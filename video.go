package twitcher

import (
	"time"
)

// Video contains meta information about a single Twitch video. More information at https://dev.twitch.tv/docs/api/reference#get-videos
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

// videosResponse represents a response from Twitch containing an array of Videos
type videosResponse struct {
	Data       []Video `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

// GetVideoByID gets a single Twitch video
func (c Client) GetVideoByID(id string) (vid Video, err error) {
	return
}

// GetVideosByIDs gets an array of Twitch videos by their IDs
func (c Client) GetVideosByIDs(ids []string) (vid []Video, err error) {
	return
}

// GetVideosByUser gets an array of Twitch videos by user ID
func (c Client) GetVideosByUser(id string) (vid []Video, err error) {
	return
}

// GetVideosByGame gets an array of Twitch videos by game ID
func (c Client) GetVideosByGame(id string) (vid []Video, err error) {
	return
}
