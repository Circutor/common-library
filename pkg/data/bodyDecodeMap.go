// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"encoding/json"
	"fmt"
)

// BodyDecodeToMap method decode body of the response object map[string]interface{}.
func (d *Data) BodyDecodeToMap(contentBody []byte) (map[string]interface{}, error) {
	var body map[string]interface{}

	err := json.Unmarshal(contentBody, &body)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", errDecodeResponseBody, err)
	}

	return body, nil
}
