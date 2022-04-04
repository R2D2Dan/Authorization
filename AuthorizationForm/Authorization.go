package authorizationform

import (
	"html/template"
	"log"
	"net/http"
)

type Users struct {
	Login    string
	Password string
}

func Authorization(w http.ResponseWriter, r *http.Request) {

	//Запоминаем данные для пользователя
	/*
	user := Users{
		Login:    r.FormValue("login"),
		Password: r.FormValue("password"),
	}
	*////

	//Подключаем html
	html, err := template.ParseFiles("./form/index.html")

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
	///

}
