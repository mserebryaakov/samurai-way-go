package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mserebryaakov/samurai-way-go/data"
)

func DialogsHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	d := data.DialogsStore
	e := json.NewEncoder(rw)
	err := e.Encode(d)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
