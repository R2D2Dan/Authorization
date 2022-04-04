package authorizationform

import (
	"html/template"
	"log"
	"net/http"
)

func Authorization(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./AuthorizationForm/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Eror Server: read file html", 500)
		return
	}
	err = html.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Eror Server: read file html", 500)
		return
	}

}
