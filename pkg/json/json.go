// Copyright (c) 2021 Circutor S.A. All rights reserved.

package json

import (
	"encoding/json"
	"net/http"
)

const (
	// Constants to use with the content type.
	contentType     = "Content-type"
	contentTypeJSON = "application/json"
)

// SendJSON convert type interface to json. Method common.
func SendJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set(contentType, contentTypeJSON)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
