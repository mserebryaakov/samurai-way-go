package data

import (
	"encoding/json"
	"io"
)

func (s *DialogPage) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(s)
}

func (s *DialogPage) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}
