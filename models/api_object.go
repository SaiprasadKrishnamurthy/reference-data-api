package models

import (
	"encoding/json"
	"io"
)

// APIObject the base API object.
type APIObject struct {
}

// ToJSON converts to json.
func (a *APIObject) ToJSON(obj interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(obj)

}
