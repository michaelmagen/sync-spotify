package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/michaelmagen/sync-spotify/config"
)

func AuthRoute(r chi.Router) {
	r.Get("/", exchangeAuthorizationCode)
	// TODO: Create route for refresh token
}

// Exchanges authorization token for access and refresh token
// Puts them in cookies and redirects back to front end
func exchangeAuthorizationCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	// Exchange the authorization code for an access token
	token, err := config.NewSpotifyOauthConfig().Exchange(context.Background(), code)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Failed to exchange code: %v", err), http.StatusInternalServerError)
		return
	}
	// Set the access token in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "access_token",
		Value: token.AccessToken,
		Path:  "/",
	})
	// Convert expiration time, to duration until experation in seconds
	expiresIn := token.Expiry.Sub(time.Now()).Seconds()

	// Set the expiry time in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "token_expiry",
		Value: fmt.Sprintf("%d", int(expiresIn)),
		Path:  "/",
	})

	// Set the refresh token in a cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "refresh_token",
		Value: token.RefreshToken,
		Path:  "/",
	})
	// Redirect the user to the FrontendURL
	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}
