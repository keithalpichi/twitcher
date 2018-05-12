package twitcher

import (
	"testing"
	"time"
)

const (
	testTime = int64(1257894000)
)

func TestTokenValidity(t *testing.T) {
	cases := []struct {
		token    Token
		expected bool
		msg      string
	}{
		{
			token: Token{
				AccessToken: "valid",
				Exp:         testTime,
			},
			expected: false,
			msg:      "Invalid token date",
		},
		{
			token: Token{
				Date:        time.Unix(testTime, 0),
				AccessToken: "",
				Exp:         testTime,
			},
			expected: false,
			msg:      "Invalid token access token",
		},
		{
			token: Token{
				Date:        time.Unix(testTime, 0),
				AccessToken: "valid",
				Exp:         0,
			},
			expected: false,
			msg:      "Invalid expiry",
		},
	}
	for _, c := range cases {
		if got := c.token.valid(); got != c.expected {
			t.Errorf("%s. Got %t, expected %t", c.msg, got, c.expected)
		}
	}
}

func TestTokenExpiry(t *testing.T) {
	token := Token{
		Date: time.Unix(testTime, 0),
		Exp:  testTime,
	}
	if got := token.Expiry(); got != time.Unix(testTime, 0) {
		t.Errorf("Got %v, expected %v", got, token.Exp)
	}
}

func TestTokenExpiresAfter(t *testing.T) {
	cases := []struct {
		token    Token
		ts       time.Time
		expected bool
		msg      string
	}{
		{
			token: Token{
				Exp: testTime,
			},
			ts:       time.Unix(testTime, 0).AddDate(0, 0, 1),
			expected: true,
			msg:      "Token should have expired",
		},
		{
			token: Token{
				Exp: testTime,
			},
			ts:       time.Unix(testTime, 0).AddDate(0, 0, -1),
			expected: false,
			msg:      "Token should not have expired",
		},
	}
	for _, c := range cases {
		if got := c.token.ExpiresAfter(c.ts); got != c.expected {
			t.Errorf("%s. Got %t, expected %t", c.msg, got, c.expected)
		}
	}
}

func TestTokenExpiresIn(t *testing.T) {
	sampleDuration, _ := time.ParseDuration("24h")
	cases := []struct {
		token    Token
		dur      time.Duration
		expected bool
		msg      string
	}{
		{
			token: Token{
				Exp: time.Now().AddDate(0, 0, 2).Unix(),
			},
			dur:      sampleDuration,
			expected: false,
			msg:      "Token should not have expired",
		},
		{
			token: Token{
				Exp: time.Now().Unix(),
			},
			dur:      sampleDuration,
			expected: true,
			msg:      "Token should have expired",
		},
	}
	for _, c := range cases {
		if got := c.token.ExpiresIn(c.dur); got != c.expected {
			t.Errorf("%s. Got %t, expected %t", c.msg, got, c.expected)
		}
	}
}

func TestTokenAuthHeader(t *testing.T) {
	token := Token{AccessToken: "1234abcd"}
	expected := "Bearer 1234abcd"
	if got := token.AuthorizationHeader(); got != expected {
		t.Errorf("Got '%s', expected '%s'", got, expected)
	}
}
