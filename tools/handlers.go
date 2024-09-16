package tools

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	data PageData
	tmp  *template.Template
)

func init() {
	tmp = template.Must(template.ParseGlob("templates/*.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		http.Error(w, "this page is not found", 404)
		return
	}
	apiURL := "https://groupietrackers.herokuapp.com/api"
	cards, err := FetchArtistData(apiURL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error fetching artist data: %v", err)
		return
	}

	artistNames, members, creationDates, firstAlbums, locations := uniqueValues(cards)

	data = PageData{
		Cards:         cards,
		ArtistNames:   artistNames,
		Members:       members,
		CreationDates: creationDates,
		FirstAlbums:   firstAlbums,
		Locations:     locations,
	}

	if err := tmp.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func uniqueValues(cards []Card) (map[string]bool, map[string]bool, map[string]bool, map[string]bool, map[string]bool) {
	artistNames := make(map[string]bool)
	members := make(map[string]bool)
	creationDates := make(map[string]bool)
	firstAlbums := make(map[string]bool)
	locations := make(map[string]bool)

	for _, card := range cards {
		artistNames[card.Name] = true
		for _, member := range card.Members {
			members[member] = true
		}
		creationDates[strconv.Itoa(card.CreationDate)] = true
		firstAlbums[card.FirstAlbum] = true
		for _, location := range card.Locations {
			locations[location] = true
		}
	}

	return artistNames, members, creationDates, firstAlbums, locations
}

func SearchResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only GET method is allowed", http.StatusBadRequest)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Bad Request: No search query provided", http.StatusBadRequest)
		return
	}

	query = strings.ToLower(query)
	var results []Card

	for _, card := range data.Cards {
		// Check for name match
		if strings.Contains(strings.ToLower(card.Name), query) {
			results = append(results, card)
		} else if Containesitem(card.Members, query) {
			results = append(results, card)
		} else if strings.Contains(strings.ToLower(card.FirstAlbum), query) {
			results = append(results, card)
		} else if strconv.Itoa(card.CreationDate) == query {
			results = append(results, card)
		} else if Containesitem(card.Locations, query) {
			results = append(results, card)
		}
	}

	// Render the results to the template
	if err := tmp.ExecuteTemplate(w, "search.html", results); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func Containesitem(slice []string, query string) bool {
	for _, item := range slice {
		if strings.Contains(strings.ToLower(item), query) {
			return true
		}
	}
	return false
}

func Bandinfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only GET method is allowed", http.StatusBadRequest)
	}
	id := strings.TrimPrefix(r.URL.RawQuery, "=id")
	if id == "" {
		http.Error(w, "/400", http.StatusBadRequest)
		return
	}
	uno, err := strconv.Atoi(id)
	if err != nil || uno > 52 || uno < 0 {
		http.Error(w, "this page is not found", 404)
		return
	}
	tmp.ExecuteTemplate(w, "bandsinfo.html", data.Cards[uno-1])
}
