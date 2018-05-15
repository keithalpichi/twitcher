# Twitcher

## What is Twitcher?
A Go package to interact with Twitch's API.

## Installation
`go get -u github.com/keithalpichi/twitcher`

## Usage
First head to [Twitch](https://dev.twitch.tv/docs/authentication/#registration) to create an application to get a Client ID and Secret

Create a Twitcher client with a `twitcher.AppConfig` and `twitcher.Client`
```
config := twitcher.AppConfig{
  ID: "your-twitch-client-id",
  Secret: "your-twitch-client-secret",
}
c := twitcher.Client(config)
```

Get an app access token from Twitch.
```
err := c.AppAccessToken()
```

If it was successful, the token `twitcher.Token` will be set on the client. To get the token access it at:
```
c.AccessToken
```
For all subsequent API requests you can:

**1.**
Utilize all the `twitcher.Client` request functions. Here are two:

Get Twitch user by ID
```
user, err := c.UserByID("123456")
```
Get Twitch users by Login (username)
```
users, err := c.UsersByLogin([]string{"user1", "user2"})
```
See the rest in the API Reference section below


**2.**
Handle the request and response yourself using `twitcher.Request` and the `twitcher.Client.Request` function:
```
v := url.Values{}
v.Set("login", "a-twitch-user")
opts := twitcher.Request{
  HTTP: http.Request{
    Method: http.MethodGet,
    Form: v,
  },
  URI: twitcher.EndPointUsers,
}
resp, err := c.Request(opts)
```

# API Reference

### Users
|Description|Function|
|:---|:---|
|User by ID|`UserByID(string) twitcher.User, error`|
|User by Login|`UserByLogin(string) twitcher.User, error`|
|Users by ID|`UsersByID([]string) []twitcher.User, error`|
|Users by Login|`UserByLogin([]string) []twitcher.User, error`|

### Videos
|Description|Function|
|:---|:---|
|Videos by User|`VideosByUser(string) []twitcher.Video, error`|
|Videos by Game ID|`VideosByGameID(string) []twitcher.Video, error`|
|Video by ID|`VideoByID(string) twitcher.Video, error`|
|Videos by ID|`VideosByID([]string) []twitcher.Video, error`|

### User Follows
|Description|Function|
|:---|:---|
|User Followers|`UsersFollowing(string) []twitcher.User, error`|
|User Following|`FollowedByUser(string) []twitcher.User, error`|
