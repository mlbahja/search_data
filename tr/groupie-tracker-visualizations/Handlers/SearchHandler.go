package groupie

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"groupie/global"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch the artist data from the API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var artists []global.ArtistData
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		http.Error(w, "Error decoding artist data", http.StatusInternalServerError)
		return
	}

	// Get the search query from the request
	query := strings.ToLower(r.URL.Query().Get("id"))
	var results []global.SearchResult

	// Perform the search
	for _, artist := range artists {
		// Search by artist/band name
		if strings.Contains(strings.ToLower(artist.Name), query) {
			results = append(results, global.SearchResult{
				ID:   artist.ID,
				Name: artist.Name,
				Type: "artist/band",
			})
		}

		// Search by members
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				results = append(results, global.SearchResult{
					ID:   artist.ID,
					Name: member,
					Type: "member",
				})
			}
		}

		// Search by creation date (convert int to string for comparison)
		creationDate := strconv.Itoa(artist.CreationDate)
		if strings.Contains(creationDate, query) {
			results = append(results, global.SearchResult{
				ID:   artist.ID,
				Name: artist.Name,
				Type: "creation date",
			})
		}

		// Search by first album date
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			results = append(results, global.SearchResult{
				ID:   artist.ID,
				Name: artist.FirstAlbum,
				Type: "first album",
			})
		}
	}

	// Return search results as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
