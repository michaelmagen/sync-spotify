package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/michaelmagen/sync-spotify/config"
	"golang.org/x/oauth2"
)

func EnsureAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the access_code from the request headers
		accessCode := r.Header.Get("access_code")

		if accessCode == "" {
			log.Println("No access code in request")
			http.Error(w, "No access code provided", http.StatusUnauthorized)
			return
		}
		// Get the token json from the Redis Store
		tokenJSON, err := config.RedisClient.Get(r.Context(), accessCode).Result()
		if err != nil {
			log.Println("Access token not in redis store")
			http.Error(w, "Token not found", http.StatusUnauthorized)
			return
		}

		// Convert JSON into token object
		var token oauth2.Token
		err = json.Unmarshal([]byte(tokenJSON), &token)
		if err != nil {
			log.Println("Failed to unmarshall the json token")
			http.Error(w, "Could not unmarshall token", http.StatusUnauthorized)
			return
		}

		// If the token is no longer valid, attempt to refresh the token
		if !token.Valid() {
			// Create token source that can refresh tokens
			tokenSource := config.OauthConfig.TokenSource(r.Context(), &token)

			// Attempt to refresh the token
			refreshedToken, err := tokenSource.Token()
			if err != nil {
				log.Println("Failed to refresh token:", err)
				http.Error(w, "Failed to refresh token", http.StatusUnauthorized)
				return
			}

			// Add new token to redis
			err = config.RedisClient.SetToken(r.Context(), *refreshedToken)
			if err != nil {
				log.Println("Failed to marshal refreshed token:", err)
				http.Error(w, "Failed to refresh token", http.StatusUnauthorized)
				return
			}
			log.Println("Adding to redis accessToken:", refreshedToken.AccessToken)

			// Delete the token in the Redis Store
			err = config.RedisClient.Del(r.Context(), accessCode).Err()
			if err != nil {
				log.Println("could not delete from redis:", err)
			}

			log.Println("removing from redis accessCode:", accessCode)

			// Set the token var to the new token
			token = *refreshedToken
		}

		ctx := context.WithValue(r.Context(), "accessToken", token.AccessToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
