// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data_test

import (
	"testing"

	"github.com/circutor/common-library/pkg/data"
	"github.com/stretchr/testify/assert"
)

func TestErrorDeviceIdDecode(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	uuid, err := d.DeviceIDDecode("*")

	assert.Equal(t, 0, len(uuid))
	assert.Contains(t, err.Error(), "error in deviceID Decode")
}

func TestSuccessErrorDeviceIdDecode(t *testing.T) {
	t.Parallel()

	d := data.NewData()

	uuid, _ := d.DeviceIDDecode("Kp5JHdH8q3LVrXqutMo2G5")
	assert.Equal(t, 36, len(uuid))
	assert.Equal(t, "123e4567-e89b-12d3-a456-426655440000", uuid)
}
