// Copyright (c) 2021 Circutor S.A. All rights reserved.

package vertical

import (
	"fmt"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/thingsboard-methods/pkg/controller"
	"github.com/sigurn/crc8"
)

// CreateCRC8 create crc8 for a DeviceId.
func CreateCRC8(deviceID string) string {
	crcUUID := fmt.Sprintf("%x", crc8.Checksum([]byte(deviceID), crc8.MakeTable(crc8.CRC8_MAXIM)))

	if len(crcUUID) == 1 {
		crcUUID = "0" + crcUUID
	}

	return crcUUID
}

// InterfaceVertical Operations device vertical.
type InterfaceVertical interface {
	SetAttrServer(deviceID, costumerID, secretKey, token, msg string, expirationTime int,
		tb controller.ThingsBoardController, data data.InterfaceData) (int, map[string]interface{}, error)
	SetAttrClient(accessToken, msg string, tb controller.ThingsBoardController,
		data data.InterfaceData) (int, map[string]interface{}, error)
	UpdateTypology(deviceID, token, msg string, tb controller.ThingsBoardController,
		data data.InterfaceData) (int, map[string]interface{}, error)
	TelemetryInitialization(accessToken, deviceType, msg string,
		tb controller.ThingsBoardController) (int, map[string]interface{}, error)
}

//go:generate mockery --name InterfaceVertical --structname InterfaceVerticalMock --filename InterfaceVerticalMock.go
