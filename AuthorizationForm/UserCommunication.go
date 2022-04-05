package authorizationform

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Login    string
	Password string
}
type Data struct {
	DataBaseName string
	Users        []User
}

//Тип база
func DB() Data {
	DataDB := Data{
		DataBaseName: "mssql: User",
		Users: []User{
			User{Login: "R2D2Dan", Password: "qwerty1"},
			User{Login: "ArtLebedev", Password: "YaKryt3000"},
			User{Login: "Vasya2020", Password: "test"},
			User{Login: "test@test.com", Password: "test"},
		},
	}
	return DataDB

}

func Authorization(w http.ResponseWriter, r *http.Request) {

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

func UserAuthorization(w http.ResponseWriter, r *http.Request) {

	user := r.FormValue("user")
	pass := r.FormValue("pass")

	//Если пользователь не ввел данные отправляем обратно... пока так.
	//Хрень но пока не знаю как по другому
	if user == "" || pass == "" {
		http.Redirect(w, r, "/Authorization", http.StatusSeeOther)

	}

	if func(u, p string, db Data) bool {

		for _, us := range db.Users {
			if us.Login == u && us.Password == p {
				return true
			}
		}
		return true
	}(user, pass, DB()) {
		fmt.Println("Пользователь найден:")
		fmt.Printf("\tUser:%s\n\tPassword:%s\n", user, pass)
	} else {
		fmt.Println("Пользователь не найден")
		http.Redirect(w, r, "/Authorization", http.StatusSeeOther)
	}

}
