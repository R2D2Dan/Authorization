package main

import (
	form "Authorization/Forms"
	"log"
	"net/http"
)

func main() {
	start()
}

func start() {

	servmux := http.NewServeMux()
	//Форма авторизации
	servmux.HandleFunc("/", form.Authorization)
	servmux.HandleFunc("/User", form.User)

	//Обработчик файлов(стилей)
	fileserver := http.FileServer(http.Dir("./templates/"))
	servmux.Handle("/templates/", http.StripPrefix("/templates/", fileserver))

	log.Println("Start Server")
	err := http.ListenAndServe(":8081", servmux)
	if err != nil {
		log.Println("error server:", err)
	}
}
