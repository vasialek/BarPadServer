package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Decode decode JSON to object
func Decode(rq *http.Request, data interface{}) error {
	ba, err := ioutil.ReadAll(rq.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(ba, data)
}

// DecodeAndDump decodes JSON to object and dump request
func DecodeAndDump(rq *http.Request, data interface{}) error {
	ba, err := ioutil.ReadAll(rq.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(ba))
	return json.Unmarshal(ba, data)
}
