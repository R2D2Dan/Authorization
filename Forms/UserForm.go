package www

import (
	database "Authorization/DataBase"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type UserData struct {
	FIO         string
	Age         string
	NumberPhone string
	Rating      string
	Score       string
}

func User(w http.ResponseWriter, r *http.Request) {

	login := r.FormValue("user")
	pass := r.FormValue("pass")

	if login == "" && pass == "" {
		log.Println("Нет данных")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id_clinet, user, ok := database.IsClient(login, pass)

	if !ok {

		log.Println("Пользователь не найден")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		log.Println(id_clinet)
		log.Println(user)
		log.Println("Пользователь найден")

	}

	u := UserData{
		FIO:         fmt.Sprintf("%s %s", user.User[id_clinet]["FirtName"], user.User[id_clinet]["LastName"]),
		Age:         fmt.Sprintf("%s", user.User[id_clinet]["Age"]),
		NumberPhone: fmt.Sprintf("%s", user.User[id_clinet]["NumberPhone"]),
		Rating:      fmt.Sprintf("%s", user.User[id_clinet]["Raiting"]),
		Score:       fmt.Sprintf("%s", user.User[id_clinet]["Score"]),
	}

	//Подключаем html
	html, err := template.ParseFiles("./templates/User/index.html")

	if err != nil {
		log.Println(err)
		http.Error(w, "Eror Server: read file html", 500)
		return
	}

	err = html.Execute(w, u)
	if err != nil {
		log.Println(err)
		http.Error(w, "Eror Server: read file html", 500)
		return
	}

}
