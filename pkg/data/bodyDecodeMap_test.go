// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data_test

import (
	"testing"

	"github.com/circutor/common-library/pkg/data"
	"github.com/stretchr/testify/assert"
)

func TestErrorDecodeMap(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	var v map[string]interface{}

	data, err := d.BodyDecodeToMap([]byte(""))

	assert.Equal(t, 0, len(data))
	assert.Equal(t, v, data)
	assert.Contains(t, err.Error(), "error in decode response body")

	data, err = d.BodyDecodeToMap(nil)

	assert.Equal(t, 0, len(data))
	assert.Equal(t, v, data)
	assert.Contains(t, err.Error(), "error in decode response body")
}

func TestSuccessDecodeMap(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	data, _ := d.BodyDecodeToMap([]byte(`{"page": 1, "fruits": ["apple", "peach"]}`))
	assert.Equal(t, 2, len(data))
}
