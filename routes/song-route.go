package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/michaelmagen/sync-spotify/middleware"
)

func SongRoute(r chi.Router) {
	// TODO: Create route for refresh token
	r.With(middleware.EnsureAuthorization).Get("/", searchSpotify)
}

func searchSpotify(w http.ResponseWriter, r *http.Request) {
	// Create a simple JSON response
	accessCode, ok := r.Context().Value("accessToken").(string)
	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("gettign accesscode from context:", accessCode)
	response := map[string]string{"message": "hello", "code": accessCode}

	// Convert the response to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonResponse)
}
