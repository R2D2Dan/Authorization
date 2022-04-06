package www

import (
	database "Authorization/DataBase"
	"log"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {

	user := r.FormValue("user")
	pass := r.FormValue("pass")

	if user == "" && pass == "" {
		log.Println("Нет данных")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if func(u, p string, db database.Data) bool {
		for _, k := range db.Users {
			if k.Login == u && k.Password == p {
				return true
			}
		}
		return false
	}(user, pass, database.DB()) {
		log.Println("Пользователь найден")
	} else {
		log.Println("Пользователь не найден")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

}
