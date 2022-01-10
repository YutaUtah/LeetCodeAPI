// this unmarshal JSON format data

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// r is whatever we get from user. x is the interface
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			fmt.Println("error in unmarshalJSON")
			return
		}
	}
}
