package funcs

import (
	"encoding/json"
	"net/http"
	"strings"
)

func ArtistSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		RenderTemplate(w, "error.html", MethodNotAllowed)
		return
	}

	// Fetch all artists
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		InternalServerError.Message = "Failed to fetch artists."
		RenderTemplate(w, "error.html", InternalServerError)
		return
	}
	defer resp.Body.Close()

	var artists []Artists
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		InternalServerError.Message = "Failed to parse artist data."
		RenderTemplate(w, "error.html", InternalServerError)
		return
	}

	// Take form relations API
	locResp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil || locResp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusNotFound)
		RenderTemplate(w, "error.html", NotFound)
		return
	}
	defer locResp.Body.Close()

	var location Locations
	if err := json.NewDecoder(locResp.Body).Decode(&location); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		InternalServerError.Message = "Failed to fetch Locations."
		RenderTemplate(w, "error.html", InternalServerError)
		return
	}

	data := PageData{
		Artist:   artists,
		Location: location,
	}

	// Get search query from URL
	query := r.URL.Query().Get("searchQuary")
	// Call search function
	filtered := SearchArtists(query, data)

	// When nothing match user quary search
	if len(filtered) == 0 {
		w.WriteHeader(http.StatusNotFound) // Look at later
	}

	// Remove unwanted objects
	filtered = FilterArtists(filtered, ExcludeIDs)
	data.Artist = filtered

	RenderTemplate(w, "index.html", data)
}

func HandleQueries(w http.ResponseWriter, r *http.Request) {
	// Check for searchQuery URL parameter
	query, ok := r.URL.Query()["searchQuary"]
	if ok && len(query) > 0 && strings.TrimSpace(query[0]) != "" {
		ArtistSearchHandler(w, r) //  Valid query
		return
	} else if ok && len(query) > 0 && strings.TrimSpace(query[0]) == "" {
		w.WriteHeader(http.StatusBadRequest) // searchQuary is present but empty || "Will look at later"
		Handler(w, r)
		return
	}

	Handler(w, r) // No searchQuary param at all: (load normal home page)
}
