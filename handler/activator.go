package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mserebryaakov/samurai-way-go/data"
)

func ActivatorHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	var userKey UserKey

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("body read")
		http.Error(rw, "Error request (read body)", http.StatusInternalServerError)
	} else {
		err = json.Unmarshal(body, &userKey)
		if err != nil {
			fmt.Println("can't unmarshal")
			http.Error(rw, "Error request (unmarshal)", http.StatusInternalServerError)
		}
	}

	fmt.Println("add key")
	data.AddKey(userKey.Key)
}
