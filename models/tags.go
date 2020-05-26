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

// FromJSON - encodes the object to JSON
func (t *Tags) FromJSON(r io.Reader) error {
	return t.APIObject.FromJSON(t, r)
}
