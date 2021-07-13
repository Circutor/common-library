// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"encoding/json"
	"net/http"
)

// SendJSON convert type interface to json. Method common.
func (d *Data) SendJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set(contentType, contentTypeJSON)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
