package twitcher

import (
	"fmt"
	"time"
)

// Token represents a twitch token object returned from a token request
type Token struct {
	Date         time.Time
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	Exp          int64    `json:"expires_in"`
	Scope        []string `json:"scope"`
}

// AuthorizationHeader returns the string "Bearer " with `t.AccessToken` appended to it
func (t Token) AuthorizationHeader() (s string) {
	s = fmt.Sprintf("Bearer %s", t.AccessToken)
	return
}

// Expiry returns the time.Time of the token expiration.
func (t Token) Expiry() (ts time.Time) {
	return time.Unix(t.Exp, 0)
}

// IsExpired returns true if the token expiry is after the time when this function is called, false otherwise
func (t Token) IsExpired() (b bool) {
	return t.ExpiresAfter(time.Now())
}

// ExpiresAfter returns true if the token expiry is after `ts` time.Time, false otherwise
func (t Token) ExpiresAfter(ts time.Time) (b bool) {
	return ts.After(t.Expiry())
}

// ExpiresIn returns true if the token expiry is after the current time plus the `dur` time.Duration, false otherwise
func (t Token) ExpiresIn(dur time.Duration) (b bool) {
	return t.ExpiresAfter(time.Now().Add(dur))
}

func (t Token) valid() (b bool) {
	if (t.Date == time.Time{} || t.AccessToken == "" || t.Exp == int64(0)) {
		return
	}
	return true
}
