package config

import (
	"golang.org/x/oauth2"
)

const (
	SpotifyRedirectURI = "http://localhost:3000/auth"
)

var spotifyEndpoint = oauth2.Endpoint{
	AuthURL:  "https://accounts.spotify.com/authorize",
	TokenURL: "https://accounts.spotify.com/api/token",
}

func NewSpotifyOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     getEnv("SPOTIFY_CLIENT_ID", "word"),
		ClientSecret: getEnv("SPOTIFY_CLIENT_SECRET", ""),
		RedirectURL:  SpotifyRedirectURI,
		Scopes:       []string{"user-read-email", "user-read-private", "streaming", "user-library-modify", "user-library-read"}, // Add required scopes
		Endpoint:     spotifyEndpoint,
	}
}
