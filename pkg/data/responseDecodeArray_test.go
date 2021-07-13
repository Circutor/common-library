// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data_test

import (
	"testing"

	"github.com/circutor/common-library/pkg/data"
	"github.com/go-playground/assert/v2"
)

func TestReturnArrayEmpty(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	var v []interface{}

	data, _ := d.ResponseDecodeToArray(nil)
	assert.Equal(t, 0, len(data))
	assert.Equal(t, v, data)

	data, _ = d.ResponseDecodeToArray("")
	assert.Equal(t, 0, len(data))
	assert.Equal(t, v, data)
}

func TestForceErrorArray(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	data, _ := d.ResponseDecodeToArray(1)
	assert.Equal(t, 0, len(data))

	data, _ = d.ResponseDecodeToArray(map[string]interface{}{"foo": make(chan int)})
	assert.Equal(t, 0, len(data))
}
