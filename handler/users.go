package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mserebryaakov/samurai-way-go/data"
)

func UsersHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	d := data.UsersStore
	e := json.NewEncoder(rw)
	err := e.Encode(d)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
