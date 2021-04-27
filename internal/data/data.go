// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	// Constants error.
	errEncodeRequestBody  = "error in encode request body"
	errDecodeResponseBody = "error in decode response body"
)

// BodyEncode method encode request body payload.
func BodyEncode(v interface{}) (*bytes.Buffer, error) {
	reqBody := new(bytes.Buffer)

	err := json.NewEncoder(reqBody).Encode(v)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", errEncodeRequestBody, err)
	}

	return reqBody, nil
}

// BodyDecode method decode body of the response object.
func BodyDecode(contentBody []byte) (map[string]interface{}, error) {
	var body map[string]interface{}

	err := json.Unmarshal(contentBody, &body)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", errDecodeResponseBody, err)
	}

	return body, nil
}

// ResponseDecode method transforms struct to map[string]interface{}.
func ResponseDecode(v interface{}) (map[string]interface{}, error) {
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
