package www

import (
	"html/template"
	"log"
	"net/http"
)

func Authorization(w http.ResponseWriter, r *http.Request) {

	//Подключаем html
	html, err := template.ParseFiles("./templates/AuthForm/Auth.html")

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


