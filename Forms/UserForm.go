package www

import (
	_ "Authorization/DataBase"
	"log"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {

	user := r.FormValue("user")
	pass := r.FormValue("pass")

	if user == "" && pass == "" {
		log.Println("Нет данных")
		http.Redirect(w, r, "/", http.StatusSeeOther)

	}

}
