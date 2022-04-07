package www

import (
	database "Authorization/DataBase"
	"fmt"
	"log"
	"net/http"
)

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
	w.Write([]byte(fmt.Sprintf("%s", user.User[id_clinet]["Login"])))

}
