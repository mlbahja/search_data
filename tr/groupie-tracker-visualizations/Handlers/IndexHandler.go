package groupie

import (
	"encoding/json"
	"net/http"
	"text/template"

	"groupie/global"
)

// IndexHandler handles requests to the root path and renders the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	// Fetch the artist data from the API
	var artists []global.ArtistData

	response, err := http.Get(global.Api + "/artists")
	if err != nil {
		ErrorPages(w, 500)
		return
	}

	defer response.Body.Close()

	// Decode the JSON response into the artists slice
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		ErrorPages(w, 500)
		return
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPages(w, 500)
		return
	}

	// Render the template with the artist data
	if err := tmpl.Execute(w, artists); err != nil {
		ErrorPages(w, 500)
		return
	}
}
