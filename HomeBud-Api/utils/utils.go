package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//ParseBody provides
func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		err = json.Unmarshal([]byte(body), x)
		if err != nil {
			return
		}
	}
}

//WriteResponse provides
func WriteResponse(w http.ResponseWriter, responseCode int, res []byte) {
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(responseCode)
	w.Write(res)
}
