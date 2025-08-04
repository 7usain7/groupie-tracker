package funcs

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		RenderTemplate(w, "error.html", MethodNotAllowed)
		return
	}

	// Call the API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		InternalServerError.Message = "Failed to get API data"
		RenderTemplate(w, "error.html", MethodNotAllowed)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON into struct
	var apiData []Artists
	if err := json.NewDecoder(resp.Body).Decode(&apiData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		InternalServerError.Message = "Failed to parse API response"
		RenderTemplate(w, "error.html", MethodNotAllowed)
		return
	}

	filtered := FilterArtists(apiData, ExcludeIDs)

	homeData := PageData{
		Artist: filtered,
	}

	RenderTemplate(w, "index.html", homeData)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		RenderTemplate(w, "error.html", NotFound)
		return
	}

	// Get query parameters
	queryParams := r.URL.Query()
	if len(queryParams) == 0 {
		Handler(w, r)
		return
	}

	// Check if query parameters is allowed
	for key := range r.URL.Query() {
		if !validQueries[key] {
			w.WriteHeader(http.StatusBadRequest)
			RenderTemplate(w, "error.html", BadRequest)
			return
		}
	}

	// Handle search queries if they exist
	HandleQueries(w, r)
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data any) {

	joinMembersList := template.FuncMap{
		"join": strings.Join,
	}

	t, err := template.New(tmpl).Funcs(joinMembersList).ParseFiles("templates/" + tmpl)
	if err != nil {
		// Render custom 404 page
		w.WriteHeader(http.StatusNotFound)
		t404, _ := template.ParseFiles("templates/error.html")
		t404.Execute(w, NotFound)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		// Template execution error
		w.WriteHeader(http.StatusInternalServerError)
		t500, _ := template.ParseFiles("templates/error.html")
		t500.Execute(w, InternalServerError)
		return
	}
}
