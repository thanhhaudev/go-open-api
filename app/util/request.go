package util

import (
	"encoding/json"
	"net/http"
)

// Response use to response json
func Response(w http.ResponseWriter, b any, s int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)

	err := json.NewEncoder(w).Encode(b)
	if err != nil {
		return
	}
}

func Inputs(r *http.Request, c any) error {
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		return err
	}

	return nil
}
