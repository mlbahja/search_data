package groupie

import (
	"encoding/json"
	"net/http"
	"text/template"

	link "groupie/global"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPages(w, 405)
		return
	}
	if r.URL.Path != "/" {
		ErrorPages(w, 404)
		return
	}

	// this the variable that we stor on it  all the artists informations
	var DATA []link.ArtistData

	response, err := http.Get(link.Api + "/artists")
	if err != nil {
		ErrorPages(w, 500)
		return
	}

	// this is make sure that we close the body , like if there is an open ressourse , it make sure they colsed, like connection of netwerk.
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&DATA)
	if err != nil {
		ErrorPages(w, 500)
		return
	}

	test, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPages(w, 500)
		return
	}

	if err := test.Execute(w, DATA); err != nil {
		ErrorPages(w, 500)
		return

	}
}
