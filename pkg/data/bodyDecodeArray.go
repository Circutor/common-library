// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"encoding/json"
	"fmt"
)

// BodyDecodeToArray method decode body of the response object []interface{}.
func (d *Data) BodyDecodeToArray(contentBody []byte) ([]interface{}, error) {
	var body []interface{}

	err := json.Unmarshal(contentBody, &body)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", errDecodeResponseBody, err)
	}

	return body, nil
}
