package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/michaelmagen/sync-spotify/config"
)

var ctx = context.Background()

// TODO: Replace with env variable
var frontendURL = "http://localhost:5173"

func AuthRoute(r chi.Router) {
	r.Get("/", exchangeAuthorizationCode)
	// TODO: Create route for refresh token
}

// Exchanges authorization token for access and refresh token
// Puts them in cookies and redirects back to front end
func exchangeAuthorizationCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		log.Println("Authorization code not provided")
		// Redirect to frontend login page with error in search params
		http.Redirect(w, r, frontendURL+"/login?error=auth_failed", http.StatusFound)
		return
	}

	// Exchange the authorization code for an access token
	token, err := config.NewSpotifyOauthConfig().Exchange(ctx, code)
	if err != nil {
		log.Println("Failed to exchange code:", err)
		// Redirect to frontend login page with error in search params
		http.Redirect(w, r, frontendURL+"/login?error=auth_failed", http.StatusFound)
		return
	}
	// Set the access token in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "access_token",
		Value: token.AccessToken,
		Path:  "/",
		// Cookie should expire after 5 minutes
		Expires: time.Now().Add(time.Minute * 5),
	})

	// Convert token to json
	tokenJSON, err := json.Marshal(token)
	if err != nil {
		log.Println("Failed to convert token to JSON:", err)
		// Redirect to frontend login page with error in search params
		http.Redirect(w, r, frontendURL+"/login?error=auth_failed", http.StatusFound)
		return
	}

	// Put access token in to redis store
	// In redis: access_token string -> full token json
	err = config.RedisClient.Set(ctx, token.AccessToken, tokenJSON, 0).Err()
	if err != nil {
		log.Println("Error adding token to redis:", err)
		// Redirect to frontend login page with error in search params
		http.Redirect(w, r, frontendURL+"/login?error=auth_failed", http.StatusFound)
		return
	}

	// Redirect the user to the FrontendURL
	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}
