// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data_test

import (
	"testing"

	"github.com/circutor/common-library/pkg/data"
	"github.com/stretchr/testify/assert"
)

func TestReturnMapEmpty(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	var v map[string]interface{}

	data, _ := d.ResponseDecodeToMap(nil)
	assert.Equal(t, 0, len(data))
	assert.Equal(t, v, data)

	data, _ = d.ResponseDecodeToMap("")
	assert.Equal(t, 0, len(data))
	assert.Equal(t, v, data)
}

func TestForceErrorMap(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	data, _ := d.ResponseDecodeToMap(1)
	assert.Equal(t, 0, len(data))

	data, _ = d.ResponseDecodeToMap(map[string]interface{}{"foo": make(chan int)})
	assert.Equal(t, 0, len(data))
}
