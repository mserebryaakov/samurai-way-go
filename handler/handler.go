package handler

import (
	"net/http"

	"github.com/mserebryaakov/samurai-way-go/data"
)

func UsersHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	d := data.Store
	err := d.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
