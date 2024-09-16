package groupie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"groupie/global"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Fetch artist data from API or in-memory storage
	response, err := http.Get(global.Api + "/artists")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var artists []global.ArtistData
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var artist *global.ArtistData
	for _, a := range artists {
		if fmt.Sprint(a.ID) == id {
			artist = &a
			break
		}
	}
	if artist == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/profile.html")
	
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artist)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
