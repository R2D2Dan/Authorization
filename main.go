package main

import (
	form "Authorization/AuthorizationForm"
	"log"
	"net/http"
)

func main() {

	servmux := http.NewServeMux()
	servmux.HandleFunc("/Authorization", form.Authorization)

	//fileserver := http.FileServer(http.Dir("./form/"))

	//servmux.Handle("/form/", http.StripPrefix("/form", fileserver))

	log.Println("Start Server")
	err := http.ListenAndServe(":8081", servmux)
	if err != nil {
		log.Println("error server:", err)
	}

}
