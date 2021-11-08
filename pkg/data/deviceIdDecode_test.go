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

	// nolint: paralleltest
	for _, tc := range testUUIDCode {
		tc := tc

		t.Run("Test success Decode", func(t *testing.T) {
			t.Parallel()

			d := data.NewData()

			uuid, _ := d.DeviceIDDecode(tc.shortUUID)
			assert.Equal(t, 36, len(uuid))
			assert.Equal(t, tc.uuid, uuid)
		})
	}
}
