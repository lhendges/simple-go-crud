package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"errors"
)

func FromJson(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func ToJson(x interface{}) ([]byte, error) {
	if x != nil {
		return json.Marshal(x)
	}
	return nil, errors.New("no data to serialize")
}
