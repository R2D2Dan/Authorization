package main

import (
	form "Authorization/AuthorizationForm"
	"log"
	"net/http"
)

func main() {

	servmux := http.NewServeMux()
	servmux.HandleFunc("/Authorization", form.Authorization)

	log.Println("Start Server")
	err := http.ListenAndServe(":8080", servmux)
	if err != nil {
		log.Println("error server:", err)
	}

}
