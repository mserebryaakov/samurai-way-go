package handler

import (
	"fmt"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "users")
}
