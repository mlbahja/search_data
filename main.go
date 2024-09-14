package main

import (
	"fmt"
	"net/http"

	vis "groupie/Handlers"
)

func servCss(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/assets/" {
		vis.ErrorPages(w, 404)
		return
	}
	fs := http.FileServer(http.Dir("./assets"))
	http.StripPrefix("/assets/", fs).ServeHTTP(w, r)
}

func main() {
	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/assets/", servCss)
	http.HandleFunc("/", vis.IndexHandler)
	http.HandleFunc("/artists/{id}", vis.PageHandler)
	fmt.Println("\033[32mServer started at http://127.0.0.1:8080\033[0m")
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
