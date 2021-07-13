// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data_test

import (
	"testing"

	"github.com/circutor/common-library/pkg/data"
	"github.com/stretchr/testify/assert"
)

func TestErrorDecodeArray(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	var v []interface{}

	data, err := d.BodyDecodeToArray([]byte(""))

	assert.Equal(t, 0, len(data))
	assert.Equal(t, v, data)
	assert.Contains(t, err.Error(), "error in decode response body")

	data, err = d.BodyDecodeToArray(nil)

	assert.Equal(t, 0, len(data))
	assert.Equal(t, v, data)
	assert.Contains(t, err.Error(), "error in decode response body")
}

func TestSuccessDecodeArray(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	data, _ := d.BodyDecodeToArray([]byte(`["apple", "peach"]`))
	assert.Equal(t, 2, len(data))
}
