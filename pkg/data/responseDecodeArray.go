// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ResponseDecodeToArray method transforms struct to []interface{}.
func (d *Data) ResponseDecodeToArray(v interface{}) ([]interface{}, error) {
	respBodyByte, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var respBody []interface{}

	err = json.NewDecoder(bytes.NewReader(respBodyByte)).Decode(&respBody)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return respBody, nil
}
