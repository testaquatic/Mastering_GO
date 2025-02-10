package handlers

import (
	"encoding/json"
	"io"
)

func SliceToJSON(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slice)
}
