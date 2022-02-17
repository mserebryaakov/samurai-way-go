package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mserebryaakov/samurai-way-go/handler"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/dialogs", handler.DialogsHandler)
	router.HandleFunc("/users", handler.UsersHandler)
	router.HandleFunc("/password", handler.SerialNumberHandler)
	router.HandleFunc("/activate", handler.ActivatorHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
