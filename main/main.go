package main

import (
	"fmt"
	"net/http"

	"github.com/mserebryaakov/samurai-way-go/handler"
)

func main() {
	http.HandleFunc("/users", handler.UsersHandler)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
