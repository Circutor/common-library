// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ResponseDecodeToMap method transforms struct to map[string]interface{}.
func (d *Data) ResponseDecodeToMap(v interface{}) (map[string]interface{}, error) {
	respBodyByte, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var respBody map[string]interface{}

	err = json.NewDecoder(bytes.NewReader(respBodyByte)).Decode(&respBody)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return respBody, nil
}
