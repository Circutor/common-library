// Copyright (c) 2021 Circutor S.A. All rights reserved.

package request_test

import (
	"net/http"
	"testing"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/request"
	"github.com/stretchr/testify/assert"
)

func TestErrorRequest(t *testing.T) {
	t.Parallel()

	r := request.NewRequest()

	_, status, _ := r.CreateNewRequest("/", "/", "", nil, nil)
	assert.Equal(t, status, http.StatusInternalServerError)
}

func TestRequestGet(t *testing.T) {
	t.Parallel()

	r := request.NewRequest()

	_, status, _ := r.CreateNewRequest("get", "/", "Bearer token_value", nil, nil)
	assert.Equal(t, status, http.StatusInternalServerError)
}

func TestRequestContentQuery(t *testing.T) {
	t.Parallel()

	r := request.NewRequest()

	query := map[string]interface{}{
		"status": false,
	}

	_, status, _ := r.CreateNewRequest("put", "/", "Bearer token_value", nil, query)
	assert.Equal(t, status, http.StatusInternalServerError)
}

func TestRequestContentBody(t *testing.T) {
	t.Parallel()

	r := request.NewRequest()

	d := data.NewData()
	body, _ := d.BodyEncode("")

	_, status, _ := r.CreateNewRequest("post", "/", "Bearer token_value", body, nil)
	assert.Equal(t, status, http.StatusInternalServerError)
}
