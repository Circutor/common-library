// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"bytes"
	"net/http"
)

const (
	// Constants error.
	errEncodeRequestBody  = "error in encode request body"
	errDecodeResponseBody = "error in decode response body"
	// Constants to use with the content type.
	contentType     = "Content-type"
	contentTypeJSON = "application/json"
)

type InterfaceData interface {
	BodyEncode(v interface{}) (*bytes.Buffer, error)
	BodyDecodeToMap(contentBody []byte) (map[string]interface{}, error)
	BodyDecodeToArray(contentBody []byte) ([]interface{}, error)
	ResponseDecodeToMap(v interface{}) (map[string]interface{}, error)
	ResponseDecodeToArray(v interface{}) ([]interface{}, error)
	SendJSON(w http.ResponseWriter, v interface{})
}

//go:generate mockery --name InterfaceData --structname InterfaceDataMock --filename InterfaceDataMock.go

type Data struct{}

// NewData creates a new InterfaceData.
func NewData() Data {
	return Data{}
}
