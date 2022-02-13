package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mserebryaakov/samurai-way-go/handler"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", handler.UsersHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
