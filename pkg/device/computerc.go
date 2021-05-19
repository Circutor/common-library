// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
)

// ComputerC vertical attributes device type ComputerC (ComputerC6 && ComputerC12).
type ComputerC struct {
	TargetCosPhi     float32 `json:"targetCosPhi"`
	Power            float32 `json:"power"`
	Voltage          float32 `json:"voltage"`
	RejectionFilters bool    `json:"rejectionFilters"`
}

// NewComputerC created a struct ComputerC.
func NewComputerC(targetCosPhi, power, voltage float32, rejectionFilter bool) ComputerC {
	return ComputerC{
		TargetCosPhi:     targetCosPhi,
		Power:            power,
		Voltage:          voltage,
		RejectionFilters: rejectionFilter,
	}
}

// SetAttrServer add in server scope attributes from attributes from to claiming device.
func (d ComputerC) SetAttrServer(c CommonService, microservice string) (int, map[string]interface{}, error) {
	attrBody := map[string]interface{}{
		"claimingAllowed": true,
		"claimingData": map[string]interface{}{
			"secretKey":      c.secretKey,
			"expirationTime": c.expirationTime,
		},
		"userMaster": c.userID,
	}

	status, messageErr, err := c.thingsBoard.Telemetry.SaveDeviceAttributes(
		c.deviceID, "SERVER_SCOPE", c.adminToken, attrBody)
	if err != nil {
		message, _ := data.ResponseDecode(errors.NewErrMessage(fmt.Sprintf("%v", messageErr[0])))

		return status, message, errors.WrapErrFound(err, microservice)
	}

	return status, nil, nil
}

// SetAttrClient add in client scope attributes.
func (d ComputerC) SetAttrClient(c CommonService, microservice string) (int, map[string]interface{}, error) {
	attributesBody := map[string]interface{}{
		"typology": "Electricity",
	}

	status, messageErr, err := c.thingsBoard.Telemetry.SaveDeviceAttributes(
		c.deviceID, "SHARED_SCOPE", c.adminToken, attributesBody)
	if err != nil {
		message, _ := data.ResponseDecode(errors.NewErrMessage(fmt.Sprintf("%v", messageErr[0])))

		return status, message, errors.WrapErrFound(err, microservice)
	}

	attrVertical, err := data.ResponseDecode(NewComputerC(0, 0, 0, false))
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, errors.WrapErrFound(err, microservice)
	}

	status, message, err := c.thingsBoard.DeviceAPI.PostDeviceAttributes(c.accessToken, attrVertical)
	if err != nil {
		return http.StatusInternalServerError, message, errors.WrapErrFound(err, microservice)
	}

	return status, nil, nil
}
