package main

import (
	"fmt"
	"net/http"

	help "tools/tools"
)

func main() {
	// setupRoutes()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", (http.StripPrefix("/static/", fs)))

	http.HandleFunc("/", help.Index)
	http.HandleFunc("/search", help.SearchResult)
	http.HandleFunc("/bandsinfo", help.Bandinfo)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// func protectStaticFiles(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if strings.HasPrefix(r.URL.Path, "/static/") {
// 			// Check ila kayn Referer header (kay3ni ja men page dyalna)
// 			referer := r.Header.Get("Referer")
// 			if referer == "" || !strings.Contains(referer, r.Host) {
// 				http.Error(w, "Direct access forbidden", http.StatusForbidden)
// 				return
// 			}
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
