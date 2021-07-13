// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data_test

import (
	"testing"

	"github.com/circutor/common-library/pkg/data"
	"github.com/stretchr/testify/assert"
)

func TestErrorEncode(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	data, _ := d.BodyEncode(nil)
	assert.Equal(t, "null\n", data.String())

	_, err := d.BodyEncode(map[string]interface{}{"foo": make(chan int)})
	assert.Contains(t, err.Error(), "error in encode request body")
}
