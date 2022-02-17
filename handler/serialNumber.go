package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mserebryaakov/samurai-way-go/data"
)

type ServerAnswer struct {
	Answer bool `json:"answer"`
}

type UserKey struct {
	Key string `json:"key"`
}

func SerialNumberHandler(rw http.ResponseWriter, r *http.Request) {
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

	fmt.Println("check key")
	var result ServerAnswer
	result.Answer = data.CheckKey(userKey.Key)
	e := json.NewEncoder(rw)
	e.Encode(result)
}
