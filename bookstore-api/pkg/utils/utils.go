package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads JSON from an HTTP request and fills a Go struct with it.f
func ParseBody(r *http.Request, x interface{}){ // r *http.Request → the incoming HTTP request (from client). x interface{} → any Go variable (struct, map, etc.) where parsed data will be stored
	if body, err := io.ReadAll(r.Body); err == nil { // Reads the entire request body. body → raw data from the request. err → error while reading. Only continues if no error happened.
		if err := json.Unmarshal([]byte(body), x); err != nil{ // Converts JSON data into a Go struct. json.Unmarshal: takes JSON bytes and fills the Go variable x
			return
		}
	}
}