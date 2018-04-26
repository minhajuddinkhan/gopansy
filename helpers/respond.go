package helpers

import (
	"encoding/json"
	"net/http"
)

//Respond Respond
func Respond(w http.ResponseWriter, i interface{}) {

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(i)
	if err != nil {
		panic(err)
	}

}
