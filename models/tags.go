package models

import (
	"io"
)

// Tags struct.
type Tags struct {
	APIObject
	InputText string   `json:"inputText"`
	Tags      []string `json:"tags"`
}

// ToJSON - encodes the object to JSON
func (t *Tags) ToJSON(w io.Writer) error {
	return t.APIObject.ToJSON(t, w)
}
