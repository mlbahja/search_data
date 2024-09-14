package groupie

import (
	"net/http"
	"text/template"
)

type errorType struct {
	ErrorCode string
	Message   string
}

func ErrorPages(w http.ResponseWriter, code int) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		return
	} else if code == 404 {
		w.WriteHeader(http.StatusNotFound)
		err = t.Execute(w, errorType{ErrorCode: "404", Message: "Sorry, the page you are looking for does not exist."})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else if code == 405 {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err = t.Execute(w, errorType{ErrorCode: "405", Message: "Method not allowed."})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	}else if code == 400 {
		w.WriteHeader(http.StatusBadRequest)
		err = t.Execute(w, errorType{ErrorCode: "400", Message: "Bad Request."})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		t.Execute(w, errorType{ErrorCode: "500", Message: "Internal Server Error."})
	}
}
