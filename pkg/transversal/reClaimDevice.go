// Copyright (c) 2021 Circutor S.A. All rights reserved.

package transversal

import (
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/thingsboard-methods/pkg/controller"
)

// ReClaimDevice reclaim device from user.
func (t *Transversal) ReClaimDevice(deviceName, token, msg string,
	tb controller.ThingsBoardController) (int, map[string]interface{}, error) {
	status, reClaim, err := tb.Device.ReClaimDevice(deviceName, token)
	if err != nil {
		return status, reClaim, errors.WrapErrFound(err, msg)
	}

	return status, nil, nil
}
