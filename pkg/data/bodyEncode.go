// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// BodyEncode method decode body of the response object.
func (d *Data) BodyEncode(v interface{}) (*bytes.Buffer, error) {
	reqBody := new(bytes.Buffer)

	err := json.NewEncoder(reqBody).Encode(v)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", errEncodeRequestBody, err)
	}

	return reqBody, nil
}
