package device

import (
	"fmt"
	"net/http"

	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
)

// AttrTransversal transversal attributes to send to ThingsBoard.
type AttrTransversal struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Zone        string `json:"zone"`
	Typology    string `json:"typology"`
	Use         string `json:"use"`
}

// NewAttrTransversal created a struct AttrTransversal.
func NewAttrTransversal(name, description, company, zone, typology, use string) AttrTransversal {
	return AttrTransversal{
		Name:        name,
		Description: description,
		Company:     company,
		Zone:        zone,
		Typology:    typology,
		Use:         use,
	}
}

// SetAttrTransversal add in shared scope attributes from traversal type.
func (d AttrTransversal) SetAttrTransversal(c CommonService, microservice string) (int, map[string]interface{}, error) {
	attrTransversal, err := data.ResponseDecode(NewAttrTransversal(
		d.Name, d.Description, d.Company, d.Zone, d.Typology, d.Use))
	if err != nil {
		dataError, _ := data.ResponseDecode(errors.NewErrMessage(err.Error()))

		return http.StatusInternalServerError, dataError, errors.WrapErrFound(err, microservice)
	}

	status, messageErr, err := c.ThingsBoard.Telemetry.SaveDeviceAttributes(
		c.DeviceID, "SHARED_SCOPE", c.AdminToken, attrTransversal)
	if err != nil {
		message, _ := data.ResponseDecode(errors.NewErrMessage(fmt.Sprintf("%v", messageErr[0])))

		return status, message, errors.WrapErrFound(err, microservice)
	}

	return status, nil, nil
}

// DeleteDevice delete device from user.
func (d AttrTransversal) DeleteDevice(c CommonService, microservice string) (int, map[string]interface{}, error) {
	status, message, err := c.ThingsBoard.Device.ReClaimDevice(c.DeviceName, c.UserToken)
	if err != nil {
		return status, message, errors.WrapErrFound(err, microservice)
	}

	return status, nil, nil
}
