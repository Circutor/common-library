// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"fmt"

	"github.com/lithammer/shortuuid/v3"
)

// DeviceIDDecode decode uuid in base58 to uuid.
func (d *Data) DeviceIDDecode(uuid string) (string, error) {
	uuidDecode, err := shortuuid.DefaultEncoder.Decode(uuid)
	if err != nil {
		return "", fmt.Errorf("%s : %w", errDeviceIDDecode, err)
	}

	return uuidDecode.String(), nil
}
