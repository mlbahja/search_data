package groupie

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"groupie/global"
	link "groupie/global"
)

var ApiData link.ApiOfArtist

func PageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPages(w, 405)
		return
	}

	if r.URL.Query().Encode() != "" {
		ErrorPages(w, 404)
		return
	}
	IdNmber, err := strconv.Atoi(r.PathValue("id"))
	if IdNmber <= 0 || err != nil || IdNmber > 52 {
		ErrorPages(w, 404)
		return
	}
	//-------------------------------------------------------------------------//
	response1, err := http.Get(link.Api + "/artists/" + r.PathValue("id"))
	if err != nil {
		ErrorPages(w, 502)
		return
	}
	defer response1.Body.Close() //***************//
	var stock link.ArtistData
	err = json.NewDecoder(response1.Body).Decode(&stock)
	if err != nil {
		ErrorPages(w, 500)
		return
	}
	ApiData.ArtistData = stock
	//------------------------------------------------------------------------//
	response2, err := http.Get(link.Api + "/locations/" + r.PathValue("id"))
	if err != nil {

		ErrorPages(w, 502)
		return
	}
	defer response2.Body.Close() //****************//
	var stocklocation link.Locations
	err = json.NewDecoder(response2.Body).Decode(&stocklocation)
	if err != nil {
		ErrorPages(w, 500)
		return
	}
	ApiData.Locations = stocklocation
	//-----------------------------------------------------------------------//
	response3, err := http.Get(link.Api + "/dates/" + r.PathValue("id"))
	if err != nil {
		ErrorPages(w, 502)
		return
	}
	defer response3.Body.Close() //***************//
	var stockDates link.Dates
	err = json.NewDecoder(response3.Body).Decode(&stockDates)
	if err != nil {
		ErrorPages(w, 500)
		return
	}
	ApiData.Dates = stockDates
	//------------------------------------------------------------------------//
	response4, err := http.Get(link.Api + "/relation/" + r.PathValue("id"))
	if err != nil {
		ErrorPages(w, 502)
		return
	}
	defer response4.Body.Close() //*************//
	var stockRelation link.Relations
	err = json.NewDecoder(response4.Body).Decode(&stockRelation)
	if err != nil {

		ErrorPages(w, 500)
		return
	}
	//------------------------------------------------------------------------//
	test, err := template.ParseFiles("templates/result.html")
	if err != nil {
		ErrorPages(w, 500)
		return
	}
	err = test.Execute(w, ApiData)
	if err != nil {
		ErrorPages(w, 500)
		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPages(w, 405)
		return
	}
	var results []d.SearchResult
	q := strings.TrimSpace(r.URL.Query().Get("s"))
	if q == "" {
		// Redirect to home page if query is empty
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	query := strings.ToLower(q)

	// Loop through all artists to find matching names and add them to results
	for i, artist := range artists {
		if len(results) > 16 {
			// Stop if we have reached the limit of 16 results
			break
		}

		if i == 0 {
			// On the first iteration, check if artist names start with the query
			for _, ar := range global.ArtistData  {
				artistName2 := strings.ToLower(ar.Name)
				if strings.HasPrefix(artistName2, query) {
					results = append(results, d.SearchResult{
						Image: ar.Image,
						ID:    ar.ID,
						Name:  ar.Name,
						Type:  "artist/band",
					})
				}
			}
		}

		artistName := strings.ToLower(artist.Name)
		// Check if artist names contain the query
		if !strings.HasPrefix(artistName, query) && strings.Contains(artistName, query) {
			results = append(results, d.SearchResult{
				Image: artist.Image,
				ID:    artist.ID,
				Name:  artist.Name,
				Type:  "artist/band",
			})
		}

		// Check if the query matches the artist's first album name
		if strings.HasPrefix(strings.ToLower(artist.FirstAlbum), query) {
			results = append(results, d.SearchResult{
				Image: artist.Image,
				ID:    artist.ID,
				Name:  artist.FirstAlbum,
				Type:  "FirstAlbum of " + artist.Name,
			})
		}

		C_Date := strconv.Itoa(artist.CreationDate)
		// Check if the query matches the artist's creation date
		if strings.HasPrefix(strings.ToLower(C_Date), query) {
			results = append(results, d.SearchResult{
				Image: artist.Image,
				ID:    artist.ID,
				Name:  C_Date,
				Type:  "Creation Date of " + artist.Name,
			})
		}
	}

	// Loop through all artists again to find matching members and add them to results
	for i, artist := range artists {
		if len(results) > 16 {
			// Stop if we have reached the limit of 16 results
			break
		}
		if i == 0 {
			// On the first iteration, check if any member names start with the query
			for _, ar := range artists {
				for _, member := range ar.Members {
					artistName2 := strings.ToLower(member)
					if strings.HasPrefix(artistName2, query) {
						results = append(results, d.SearchResult{
							Image: ar.Image,
							ID:    ar.ID,
							Name:  member,
							Type:  "member of " + ar.Name,
						})
					}
				}
			}
		}

		// Check if any member names contain the query
		for _, member := range artist.Members {
			if !strings.HasPrefix(strings.ToLower(member), query) && strings.Contains(strings.ToLower(member), query) {
				results = append(results, d.SearchResult{
					Image: artist.Image,
					ID:    artist.ID,
					Name:  member,
					Type:  "member of " + artist.Name,
				})
			}
		}
	}

	// Loop through all location indices to find matching locations and add them to results
	j := 0
	for _, loc := range artis.Index {
		name := artists[j].Name
		for _, lo := range loc.Locations {
			// Check if the query matches any location name
			if strings.Contains(strings.ToLower(lo), query) {
				if len(results) < 16 {
					results = append(results, d.SearchResult{
						Image: artists[j].Image,
						ID:    loc.ID,
						Name:  lo,
						Type:  "location " + name,
					})
				}
			}
		}
		j++
	}

	// Attempt to parse the search results template
	tmp, err := template.ParseFiles("template/search.html")
	if err != nil {
		ErrorPages(w, 500)
		return
	}

	// Check if results are empty and handle accordingly
	if results == nil {
		// Attempt to parse the notfound template file
		tmp1, err := template.ParseFiles("template/notfound.html")
		if err != nil {
			// If template parsing fails, handle the error
			ErrorPages(w, 500)
			return
		}

		// Execute the notfound template
		err = tmp1.Execute(w, nil)
		if err != nil {
			ErrorPages(w, 500)
			return
		}

	}

	// Execute the search results template with the results
	if err := tmp.Execute(w, results); err != nil {
		ErrorPages(w, 500)
		return
	}
}
