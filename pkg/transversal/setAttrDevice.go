// Copyright (c) 2021 Circutor S.A. All rights reserved.

package transversal

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/thingsboard-methods/pkg/controller"
)

// SetAttrDevice add in shared scope attributes from traversal type.
func (t *Transversal) SetAttrDevice(deviceID, token, msg string, tb controller.ThingsBoardController,
	data data.InterfaceData) (int, map[string]interface{}, error) {
	attrTransversal, err := data.ResponseDecodeToMap(
		NewTransversal(t.Name, t.Description, t.Company, t.Zone, t.Typology, t.Use))
	if err != nil {
		dataError, _ := data.ResponseDecodeToMap(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, errors.WrapErrFound(err, msg)
	}

	status, saveAttributes, err := tb.Telemetry.SaveDeviceAttributes(
		deviceID, "SHARED_SCOPE", token, attrTransversal)
	if err != nil {
		message, _ := data.ResponseDecodeToMap(errors.NewErrMessage(fmt.Sprintf("%v", saveAttributes[0])))

		return status, message, errors.WrapErrFound(err, msg)
	}

	return status, nil, nil
}
