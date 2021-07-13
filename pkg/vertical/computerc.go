// Copyright (c) 2021 Circutor S.A. All rights reserved.

package vertical

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/translations"
	"github.com/circutor/thingsboard-methods/pkg/controller"
	"github.com/go-playground/validator/v10"
)

// ComputerC vertical attributes device type ComputerC (ComputerC6 && ComputerC12).
type ComputerC struct {
	TargetCosPhi     float32 `json:"targetCosPhi" validate:"omitempty,gte=0.89,lte=0.99"`
	Power            float32 `json:"power" validate:"omitempty,gte=0,lte=9999.9"`
	Voltage          float32 `json:"voltage" validate:"omitempty,gte=180,lte=999"`
	RejectionFilters bool    `json:"rejectionFilters" default:"true" validate:"omitempty"`
}

// NewComputerC created a struct ComputerC implement InterfaceVertical.
func NewComputerC(targetCosPhi, power, voltage float32,
	rejectionFilter bool) ComputerC {
	return ComputerC{
		TargetCosPhi:     targetCosPhi,
		Power:            power,
		Voltage:          voltage,
		RejectionFilters: rejectionFilter,
	}
}

// IsValid validates if a ComputerC struct is valid or not.
func (c ComputerC) IsValid() ([]string, error) {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		messageError := translations.TranslateEnError(err, validate)

		return messageError, errors.WrapErrFound(err, err.Error())
	}

	return nil, nil
}

// SetAttrServer add in server scope attributes from attributes from to claiming device.
func (c ComputerC) SetAttrServer(deviceID, costumerID, secretKey, token, msg string, expirationTime int,
	tb controller.ThingsBoardController, data data.InterfaceData) (int, map[string]interface{}, error) {
	attrBody := map[string]interface{}{
		"claimingAllowed": true,
		"claimingData":    map[string]interface{}{"secretKey": secretKey, "expirationTime": expirationTime},
		"userMaster":      costumerID,
		"subscriber":      map[string]interface{}{},
	}

	status, saveAttributes, err := tb.Telemetry.SaveDeviceAttributes(deviceID, "SERVER_SCOPE", token, attrBody)
	if err != nil {
		message, _ := data.ResponseDecodeToMap(errors.NewErrMessage(fmt.Sprintf("%v", saveAttributes[0])))

		return status, message, errors.WrapErrFound(err, msg)
	}

	return status, nil, nil
}

// SetAttrClient add attributes in client scope.
func (c ComputerC) SetAttrClient(accessToken, msg string, tb controller.ThingsBoardController,
	data data.InterfaceData) (int, map[string]interface{}, error) {
	attrBody, err := data.ResponseDecodeToMap(
		NewComputerC(c.TargetCosPhi, c.Power, c.Voltage, c.RejectionFilters))
	if err != nil {
		dataError, _ := data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, errors.WrapErrFound(err, msg)
	}

	status, deviceAttributes, err := tb.DeviceAPI.PostDeviceAttributes(accessToken, attrBody)
	if err != nil {
		return http.StatusInternalServerError, deviceAttributes, errors.WrapErrFound(err, msg)
	}

	return status, nil, nil
}

// UpdateTypology update attributes typology from device type.
func (c ComputerC) UpdateTypology(deviceID, token, msg string, tb controller.ThingsBoardController,
	data data.InterfaceData) (int, map[string]interface{}, error) {
	attrBody := map[string]interface{}{
		"typology": "Electricity",
	}

	status, saveAttributes, err := tb.Telemetry.SaveDeviceAttributes(deviceID, "SHARED_SCOPE", token, attrBody)
	if err != nil {
		message, _ := data.ResponseDecodeToMap(errors.NewErrMessage(fmt.Sprintf("%v", saveAttributes[0])))

		return status, message, errors.WrapErrFound(err, msg)
	}

	return status, nil, nil
}
