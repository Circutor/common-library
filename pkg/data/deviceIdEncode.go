// Copyright (c) 2021 Circutor S.A. All rights reserved.

package data

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v3"
)

// DeviceIDEncode encode uuid to base58.
func (d *Data) DeviceIDEncode(deviceID string) (string, error) {
	uuidParse, err := uuid.Parse(deviceID)
	if err != nil {
		return "", fmt.Errorf("%s : %w", errDeviceIDEncode, err)
	}

	shortID := shortuuid.DefaultEncoder.Encode(uuidParse)

	for len(shortID) < 22 {
		shortID += "2"
	}

	return shortID, nil
}
