package groupie

import (
	"encoding/json"
	"net/http"
	"strconv"
	"text/template"

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
	ApiData.Relation = stockRelation
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
